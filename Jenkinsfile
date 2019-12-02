#!/usr/bin/env groovy
pipeline {

    environment {
    registry = "sysrex/snippetreviews"
    registryCredential = 'dockerhub'
    }

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
                    docker.withRegistry('https://hub.docker.com', 'dockerhub') {

                    def customImage = snippetreviews_go:latest
                    customImage.push()
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
    post {
        always {
            // Clean up our workspace.
            deleteDir()
        }
    }
}