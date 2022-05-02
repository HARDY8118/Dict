pipeline {
    agent any
    
    tools {nodejs "Node 16.9.1"}

    stages {
        stage('Cloning') {
            steps {
                git "https://github.com/HARDY8118/Dict.git"
                sh "cd client"
            }
        }

        stage('Install dependencies') {
            steps {
                npm "install -g yarn"
                sh "yarn install"
            }
        }

        stage('Build files') {
            steps {
                sh "yarn build" 
            }
        }

        stage('Copy') {
            steps {
                sh "cd .."
                sh "rm -r Server/static/"
                sh "mkdir Server/static"
                sh "cp -r client/dist/* Server/static/"
            }
        }

        stage('Build docker build') {
            steps {
                sh "docker build -t goserver ."
            }
        }
    }
}