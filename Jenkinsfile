#!/usr/bin/env groovy
pipeline {

    environment {
    registry = "sysrex/snippetreviews"
    registryCredential = 'dockerhub'
    }

    agent any    
    // If anything fails, the whole Pipeline stops.
    
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
                    sh 'docker push snippetreviews_go:latest'
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