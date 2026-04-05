pipeline {
    agent any
    options {
        timestamps()
    }
    stages {
        stage("Init") {
            steps {
                sh "make init"
            }
        }
        stage("Lint api") {
            steps {
                sh "make lint"
            }
        }
        stage("test api") {
            steps {
                sh "make test-api"
            }
        }
        stage("Down") {
            steps {
                sh "make docker-down-clear"
            }
        }
    }
    post {
        always {
            sh "make docker-down-clear || true"
        }
    }
}
