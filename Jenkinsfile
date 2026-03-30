pipeline {
    agent any
    options {
        timestamps()
    }
    stages {
        stage("Cleanup") {
            steps {
                script {
                    // Очищаем перед началом работы
                    sh "make docker-down-clear || true"
                }
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
