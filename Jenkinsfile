#!/usr/bin/env groovy
pipeline {
    environment {
    registry = "sysrex/snippetsreview"
    registryCredential = 'dockerhub'
    }

    agent any    
    
    stages {

        stage('Test') {   
            steps {                                          
                sh 'docker-compose build'
                sh 'docker-compose up -d'           
            }            
        }

        stage('Build') {
            steps {                 
                script {
                    docker.build registry + ":$BUILD_NUMBER"
                    }
            }
        }
        stage('Deploy Image') {
            steps{    
                script {
                    docker.withRegistry( '', registryCredential ) {
                    dockerImage.push()
                    }
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