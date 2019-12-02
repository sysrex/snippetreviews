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
                    docker.withRegistry('https://hub.docker.com/sysrex/snippetreviews', 'dockerhub') {
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
    post {
        always {
            deleteDir()
            }
        }
    }
}