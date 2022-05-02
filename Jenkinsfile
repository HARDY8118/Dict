pipeline {
    agent any
    
    tools {nodejs "Node 16.9.1"}

    stages {
        stage('Cloning') {
            steps {
                git "https://github.com/HARDY8118/Dict.git"
            }
        }

        stage('Install dependencies') {
            steps {
                yarn 'install' workspaceSubdirectory 'client'
            }
        }

        stage('Build files') {
            steps {
                yarn 'build' workspaceSubdirectory 'client'
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