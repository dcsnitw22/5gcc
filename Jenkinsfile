pipeline {
    agent{
        kubernetes {
            defaultContainer 'jnlp'
            yamlFile 'agentPod.yaml'
        }
    }
    stages {
        stage('Build') {
            // agent { node { label 'label1'} }
            steps {
                echo "Running ${env.JOB_NAME} on ${env.JENKINS_URL}"
                container('docker'){
                    echo 'Building main..'
                    sh 'go version'
                    sh 'docker ps'
                    sh 'make install'
                }
                // sh 'make install'
            }
        }
        stage('Test'){
            steps{
                container('docker'){
                    echo 'Running tests..'
                    // sh 'go version'
                    // sh 'go test -v --cover'
                }
            }
        }
        stage('Build docker image'){
            steps {
                container('docker'){
                    sh 'docker -v'
                    sh 'make install'
                    sh 'make pdusmsp-prod-image'
                }
            }
        }
        stage('Push to Nexus Docker Repo') {
            when {
                environment name: 'JOB_NAME', value: 'wipro5gc'
            }
            steps {
                container('docker'){
                    echo 'Pushing to Nexus Docker Repo...'
                    sh 'docker -v'
                    // sh 'docker info'
                    // sh 'docker build -t smfbuild:v1 ./cmd/main'
                    sh 'docker login -u admin -p wipro123 10.250.108.118:8083'
                    sh 'docker tag pdusmsp_prod:latest 10.250.108.118:8083/pdusmsp_prod:${BUILD_NUMBER}'
                    sh 'docker push 10.250.108.118:8083/pdusmsp_prod:${BUILD_NUMBER}'
                    
                    // sh 'echo BUILD_NUMBER=$BUILD_NUMBER >> spinnaker.properties'
                    // archiveArtifacts(spinnaker.properties)
                    script {
                        writeFile file: 'build.properties', text: "BUILD_NUMBER=${BUILD_NUMBER}"
                        archiveArtifacts 'build.properties'
                    }
                }
            }
        }
    }
}
