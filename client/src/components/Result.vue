<script>
import Sound from "../assets/sound.vue"

export default {
    props: ['result'],
    components: {
        SoundIcon: Sound
    },
    methods: {
        playSound(src) {
            if (src) {
                this.loading = true
                this.$refs.s.setAttribute('src', src)
                this.$refs.s.setAttribute('type', "audio/" + src.split('.').slice(-1)[0])
                this.$refs.a.load()
                this.$refs.a.play()
            }
        },
        changeConfig(e) {
            console.log(e)
        }
    },
    data() {
        return {
            playing: false,
            loading: false
        }
    },
    mounted() {
        this.$refs.a.addEventListener('ended', () => {
            this.playing = false
        })
        this.$refs.a.addEventListener('loadeddata', () => {
            this.loading = false
            this.playing = true
        })
    },
    // emits: ['configchange'],
}
</script>

<template>
    <section>
        <b-button v-b-toggle.collapse-1 variant="primary">Toggle visibility</b-button>
        <b-collapse id="collapse-1" class="mt-2" visible>
            <b-card v-for="(res, i) in result" :key="i">
                <b-alert show variant="dark">{{ res.word }}</b-alert>
                <b-accordion v-if="res.meanings" free>
                    <b-accordion-item v-for="(m, j) in res.meanings" :title="'#' + (j + 1)" :key="j">
                        <strong>
                            {{ m.partOfSpeech }}
                        </strong>
                        <b-list-group>
                            <b-list-group-item v-for="(d, k) in m.definitions" :key="k">
                                {{ d.definition }}
                                <b-badge v-for="(d, i) in m.synonyms" :key="i" variant="success">
                                    {{ s }}
                                </b-badge>
                                <br>
                                <b-badge v-for="(d, i) in m.antonyms" :key="i" variant="danger">
                                    {{ a }}
                                </b-badge>
                            </b-list-group-item>
                        </b-list-group>
                        <b-badge v-for="(s, i) in m.synonyms" :key="i" variant="success">
                            {{ s }}
                        </b-badge>
                        <br>
                        <b-badge v-for="(a, i) in m.antonyms" :key="i" variant="danger">
                            {{ a }}
                        </b-badge>
                    </b-accordion-item>
                </b-accordion>
                <br>
                <b-list-group v-if="res.phonetics">
                    <b-list-group-item v-for="(p, l) in res.phonetics" :key="l">
                        <strong>{{ p.text }}</strong>
                        <SoundIcon v-if="p.audio" @click="playSound(p.audio)" />
                        <b-button variant="primary" disabled v-if="playing">
                            <b-spinner small type="grow"></b-spinner>
                            Playing...
                        </b-button>
                        <b-button variant="primary" disabled v-else-if="loading">
                            <b-spinner small type="grow"></b-spinner>
                        </b-button>
                    </b-list-group-item>
                </b-list-group>
            </b-card>
        </b-collapse>
    </section>
    <audio ref="a">
        <source ref="s">
    </audio>
</template>

<style>
section {
    margin: 1em;
}
</style>