pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'tranquockhang181004/go-shop'
        DOCKER_TAG = 'ci-cd'
        TELEGRAM_BOT_TOKEN = "7626023144:AAHKCwYHpghRHInMs7XaOes6SssmecpIon0"
        TELEGRAM_CHAT_ID = "-1002433054280"
    }

    stages {
        stage('Clone Repository') {
            steps {
                git branch: 'master', url: 'https://github.com/tranquockhang1810/go-shop.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    docker.build("${DOCKER_IMAGE}:${DOCKER_TAG}")
                }
            }
        }

        stage('Run Tests') {
            steps {
                echo 'Running tests...'
            }
        }

        stage('Push to Docker Hub') {
            steps {
                script {
                    docker.withRegistry('https://index.docker.io/v1/', 'docker-hub-credentials') {
                        docker.image("${DOCKER_IMAGE}:${DOCKER_TAG}").push()
                    }
                }
            }
        }

        stage('Deploy Golang to DEV') {
            steps {
                echo 'Deploying to DEV...'
                sh 'docker image pull tranquockhang181004/go-shop:ci-cd'
                sh 'docker container stop golang-jenkins || echo "this container does not exist"'
                sh 'docker network create dev || echo "this network exists"'
                sh 'echo y | docker container prune '

                sh 'docker container run -d --rm --name server-golang -p 4000:3000 --network dev tranquockhang181004/go-shop:ci-cd'
            }
        }
    }

    post {
        always {
            cleanWs()
        }

        success {
            sendTelegramMessage("✅ Build #${BUILD_NUMBER} was successful! ✅")
        }

        failure {
            sendTelegramMessage("❌ Build #${BUILD_NUMBER} failed. ❌")
        }
    }
}

def sendTelegramMessage(String message) {
    sh """
    curl -s -X POST https://api.telegram.org/bot${TELEGRAM_BOT_TOKEN}/sendMessage \
    -d chat_id=${TELEGRAM_CHAT_ID} \
    -d text="${message}"
    """
}