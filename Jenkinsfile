pipeline {
    agent any
    
    tools {
        nodejs "Node 16.9.1"
        go "Go"
    }

    stages {
        stage('Cloning') {
            steps {
                git "https://github.com/HARDY8118/Dict.git"
            }
        }

        stage('Install dependencies') {
            steps {
                sh "yarn --cwd client install"
            }
        }

        stage('Build client files') {
            steps {
                dir("${env.WORKSPACE}/client"){
                    sh "yarn build"
                }
            }
        }

        stage('Copy client to build directory') {
            steps {
                sh "rm -rf build/static/"
                sh "mkdir build/static"
                sh "cp -r client/dist/* build/static/"
            }
        }

        stage('Build server files') {
            environemnt {
                GIN_MODE="release"
            }
            steps {
                dir("${env.WORKSPACE}/Server"){
                    sh "go mod tidy"
                    sh "go build -o ../build/server ."
                }
            }
        }
    }
}