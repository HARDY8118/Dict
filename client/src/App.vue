<script>
import gql from 'graphql-tag'
import Result from "./components/Result.vue";
import ApolloClient from "apollo-boost";

const apolloClient = new ApolloClient({
  uri: "/api",
});

const search = (queryWord, config) => new Promise(async (resolve, reject) => {
  // config = Object.assign({
  //   showWord: true,
  //   showPhonetics: true,
  //   showMeanings: true
  // }, config)

  try {
    const { data } = await apolloClient.query({
      query: gql`query Query($word: String!, $showWord: Boolean!, $showPhonetics: Boolean!, $showMeanings: Boolean!, $showSynonyms: Boolean!, $showAntonyms: Boolean!){
        query(word: $word){
          word @include(if: $showWord)
          phonetics @include(if: $showPhonetics){
            text
            audio
          }
          meanings @include(if: $showMeanings){
            partOfSpeech
            definitions{
              definition
              synonyms @include(if: $showSynonyms)
              antonyms @include(if: $showAntonyms)
            }
            synonyms @include(if: $showSynonyms)
            antonyms @include(if: $showAntonyms)
          }
        }
      }`,
      variables: {
        word: queryWord,
        ...config
      }
    })
    resolve(data.query)
  } catch (e) {
    console.error(e)
    reject(e)
  }
});

export default {
  data() {
    return {
      query: '',
      searching: false,
      hello: '',
      online: "",
      status: "Connecting",
      config: {
        showWord: true,
        showPhonetics: false,
        showMeanings: true,
        showSynonyms: true,
        showAntonyms: true,
      },
      result: []
    }
  },
  methods: {
    async submit(e) {
      e.preventDefault()
      if (this.query && this.query.trim()) {
        try {
          this.searching = true
          this.result = await search(this.query, this.config)
          this.searching = false
        } catch (e) {
          console.error(e)
        }
      } else {
        window.alert('Required')
      }
    },
    handleConfig(prop, value) {
      this.config[prop] = value
      this.result = []
    }
  },
  components: {
    Result
  },
  async mounted() {
    try {
      const { data } = await apolloClient.query({
        query: gql`query {
            hello
        }`
      })

      if (data && data.hello == "world") {
        this.status = "Search"
        this.online = "success"
      } else {
        this.status = "Not Connected"
        this.online = "danger"
      }
    }
    catch (e) {
      // console.error(e)
      // reject(e)
      this.status = "Not Connected"
      this.online = "danger"
    }
  }
}
</script>

<template>
  <header>

  </header>
  <main>

    <b-form @submit="submit">
      <b-input-group prepend="Word">
        <b-form-input v-model="query" required />
        <b-input-group-append>
          <b-button :variant="online" @click="submit" ref="b">{{ status }}</b-button>
        </b-input-group-append>
      </b-input-group>
    </b-form>
    <div class="controls">
      <b-form-checkbox switch size="lg" v-model="config.showMeanings"
        @change="(e) => handleConfig('showMeanings', e)">Show Meanings
      </b-form-checkbox>
      <b-form-checkbox switch size="lg" v-model="config.showPhonetics"
        @change="(e) => handleConfig('showPhonetics', e)">Show Phonetics
      </b-form-checkbox>
      <b-form-checkbox switch size="lg" v-model="config.showSynonyms"
        @change="(e) => handleConfig('showPhonetics', e)">Show Synonyms
      </b-form-checkbox>
      <b-form-checkbox switch size="lg" v-model="config.showAntonyms"
        @change="(e) => handleConfig('showPhonetics', e)">Show Antonyms
      </b-form-checkbox>
    </div>
    <Result v-if="!searching && result.length" @configchange="handleConfig" :result="result" />
    <div class="spinners-container" v-else-if="searching">
      <b-spinner type="grow" />
      <b-spinner type="grow" />
      <b-spinner type="grow" />
      <b-spinner type="grow" />
      <b-spinner type="grow" />
    </div>
  </main>
  <footer>

  </footer>
</template>

<style>
#app {
  padding: 2em;
  height: 100vh;
  display: flex;
  flex-direction: column;
}

main {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.controls {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
}

.spinners-container {
  text-align: center;
  margin: 2em;
}
</style>
