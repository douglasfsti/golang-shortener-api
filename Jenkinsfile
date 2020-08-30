pipeline {
    agent any
    tools {
        go 'go1.14'
    }
    environment {
        GO111MODULE = 'on'
        CGO_ENABLED = 0
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage('code formatter') {
            steps {
                echo 'code formatter analyzes the formatting of source code.'
                sh '''
                source_code=$(gofmt -l `find . -name '*.go' | grep -v vendor`)
                if [[ -n ${source_code} ]];then
                    echo "${source_code}"
                    echo "Please reformat your code"
                    exit 1
                fi
                '''
            }
        }
    }
    post {
        always {
            emailext body: "${currentBuild.currentResult}: Job ${env.JOB_NAME} build ${env.BUILD_NUMBER}\n More info at: ${env.BUILD_URL}",
                recipientProviders: [[$class: 'DevelopersRecipientProvider'], [$class: 'RequesterRecipientProvider']],
                to: "${params.RECIPIENTS}",
                subject: "Jenkins Build ${currentBuild.currentResult}: Job ${env.JOB_NAME}"

        }
    }
}
