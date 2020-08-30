pipeline {
    agent { docker { image 'golang' } }
    environment {
        GO111MODULE = 'on'
        CGO_ENABLED = 0
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage('version') {
            steps {
                sh 'go version'
            }
        }

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
}
