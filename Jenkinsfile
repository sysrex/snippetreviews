#!/usr/bin/env groovy
pipeline {

    agent any    
    
    stages {
        stage('Cloning Git') {
            steps {
                git 'https://github.com/sysrex/snippetreviews.git'
            }
        }

        stage('Build & Test') {   
            steps {                                           
                sh 'docker-compose build'
                sh 'docker-compose up -d'           
            }            
        }

        stage('Push') {
            steps {                 
                script {
                    docker.withRegistry('https://hub.docker.com', '0448bc08-3391-4a48-b70f-44e9348c6aaf') {
                    sh 'docker tag snippetreviews_go:latest sysrex/snippterviews_go:tag'
                    def customImage = "sysrex/snippterviews_go:tag"
                    customImage.push()
                    }
                }
            }
        }      

        stage('Deploy') {         
            environment {
                DOCKER_CREDENTIALS = credentials('dockerhub')
            }

            steps {                           
                // Use a scripted pipeline.
                script {
                    echo "This will deploy"
                }
            }
        }
        stage('Cleanup') {   
            steps {                                           
                sh 'docker-compose down'    
            }            
        }
    }
}