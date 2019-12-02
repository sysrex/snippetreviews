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
                    docker.build registry + ":$BUILD_NUMBER"
                }
            }
        }      

        stage('Deploy') {         
            environment {
                // Extract the username and password of our credentials into "DOCKER_CREDENTIALS_USR" and "DOCKER_CREDENTIALS_PSW".
                // (NOTE 1: DOCKER_CREDENTIALS will be set to "your_username:your_password".)
                // The new variables will always be YOUR_VARIABLE_NAME + _USR and _PSW.
                // (NOTE 2: You can't print credentials in the pipeline for security reasons.)
                DOCKER_CREDENTIALS = credentials('dockerhub')
            }

            steps {                           
                // Use a scripted pipeline.
                script {
                    node {
                        def app

                        stage('Clone repository') {
                            checkout scm
                        }

                        stage('Build image') {                            
                            app = docker.build("${env.DOCKER_CREDENTIALS_USR}/my-project-img")
                        }

                        stage('Push image') {  
                            // Use the Credential ID of the Docker Hub Credentials we added to Jenkins.
                            docker.withRegistry('https://registry.hub.docker.com', 'my-docker-credentials-id') {                                
                                // Push image and tag it with our build number for versioning purposes.
                                app.push("${env.BUILD_NUMBER}")                      

                                // Push the same image and tag it as the latest version (appears at the top of our version list).
                                app.push("latest")
                            }
                        }              
                    }                 
                }
            }
        }
        stage('Cleanup') {   
            // Use golang.
            agent { docker { image 'golang:latest' } }

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