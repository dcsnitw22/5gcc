pipeline {
    agent{
        kubernetes {
            defaultContainer 'docker'
            yamlFile 'agentPod.yaml'
        }
    }
    stages {
        stage ('Code Check') {
            steps {
                sh 'go version'
                sh 'make code-check'
            }
        }
        stage ('Build') {
            steps {
                // echo "Running ${env.JOB_NAME} on ${env.JENKINS_URL}"    
                echo 'Building main...'
                // sh 'docker ps'
                sh 'make build'
                sh 'make install'
            }
        }
        stage ('Test'){
            steps{
                echo 'Running tests..'
                // sh 'go version'
                // sh 'go test -v --cover'
                sh 'make ginkgo-test'
                // sh 'robot --version'
            }
        }
        stage ('SonarQube Code Analysis'){
            steps {
                script {
                    // requires SonarQube Scanner 2.8+
                    scannerHome = tool 'sonarqube-scanner'
                }
                sh 'java -version'
                withSonarQubeEnv('sonarqube-scanner') {
                    sh "${scannerHome}/bin/sonar-scanner"
                }
            }
        }
        stage ('Quality Gate'){
            steps{
                echo 'Quality gate'
                // waitForQualityGate abortPipeline: true
            }
        }
        stage ('Build docker image'){
            steps {
                sh 'docker -v'
                sh 'make pdusmsp-prod-image'
            }
        }
        stage ('Push to Nexus Docker Repo') {
            when {
                environment name: 'JOB_NAME', value: 'wipro5gc'
            }
            steps {
                echo 'Pushing to Nexus Docker Repo...'
                sh 'docker -v'
                // sh 'docker info'
                // sh 'docker build -t smfbuild:v1 ./cmd/main'
                sh 'docker login -u admin -p wipro123 10.250.108.118:8083'
                sh 'docker tag pdusmsp_prod:latest 10.250.108.118:8083/pdusmsp_prod:${BUILD_NUMBER}'
                sh 'docker push 10.250.108.118:8083/pdusmsp_prod:${BUILD_NUMBER}'
                
                //script {
                //    writeFile file: 'build.properties', text: "BUILD_NUMBER=${BUILD_NUMBER}"
                //    archiveArtifacts 'build.properties'
                //}

                sh """
                    curl --location 'http://10.250.108.120:8084/webhooks/webhook/jenkinsPdu' \
                    --header 'Content-Type: application/json' \
                    --data '{
                      "parameters": {
                        "tag": ${BUILD_NUMBER},
                        "clusterName": "main-cluster"
                      }
                    }'
                """
            }
        }


    }
}
