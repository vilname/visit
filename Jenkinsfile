pipeline {
    agent any
    options {
        timestamps()
    }
    stages {
        stage("Cleanup") {
            steps {
                sh "docker rm -f postgres_db"
                sh "docker rm -f go_app"
            }
        }
        stage("Init") {
            steps {
                sh "make init"
            }
        }
    }
    post {
        always {
            sh "make docker-down-clear || true"
        }
    }
}
