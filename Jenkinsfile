#!/usr/bin/env groovy
pipeline {

    agent any    
    
    stages {

        stage('Build & Test') {   
            steps {  
                checkout scm                                         
                sh 'docker-compose build'
                sh 'docker-compose up -d'           
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