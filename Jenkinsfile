#!/usr/bin/env groovy
pipeline {

    agent any    
    
    stages {

        stage('Build & Test') {   
            steps {                                          
                sh 'docker-compose build'
                sh 'docker-compose up -d'           
            }            
        }
        stage('Cleanup') {   
            steps {                                           
                sh 'docker-compose down'    
            }            
        }
    }
}