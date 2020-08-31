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

        stage('static source code analyzes') {
            agent { docker { image 'golangci/golangci-lint' }}
            sh 'golangci-lint --version'
            sh 'golangci-lint run ./...'
        }

        stage('unit test') {
            sh 'go test -v -coverprofile=coverage.out -covermode count > tests.out'

            // convert tests results
            sh "go get github.com/tebeka/go2xunit"
            sh "go2xunit < tests.out -output tests.xml"
            junit "tests.xml"

            // convert coverage
            sh "go get github.com/t-yuki/gocover-cobertura"
            sh "gocover-cobertura < coverage.out > coverage.xml"

            step([$class: 'CoberturaPublisher', coberturaReportFile: 'coverage.xml'])
        }

        stage('archive tests and cobertura reports') {
            archiveArtifacts '**/tests.out, **/tests.xml, **/coverage.out, **/coverage.xml, **/coverage2.xml'
        }
    }
}
