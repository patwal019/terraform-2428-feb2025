# Day 5

## Training Feedback - Kindly share your training feedback here
<pre>
https://survey.zohopublic.com/zs/jkCK83  
</pre>

## Info - Infrastructure automation test cases ( Sentinel Policy Management )
https://developer.hashicorp.com/terraform/tutorials/policy/sentinel-policy

## Info - Terraform Module vs Terraform Provider
<pre>
- Terraform Provider is developed by a Terraform Provider developer using golang programming language
- Terraform Module is written by DevOps Engineers who use the Terraform provider
- Terraform Module is written in HCL(Hashicorp Configuration Language) will file extensions *.tf
- Terraform Module is written in declarative code
- Terraform Modules are reusable code that can be invoked from Terraform Root modules
- Terraform Modules can be reused across many Terraform Projects
- Terraform Modules depends on one or more Terraform Providers
- Terraform Projects is nothing but the top level folder that has the below
  - Terraform Providers
  - Terraform Root Module
    - Terrafoorm child Module(s) optionally
</pre>

## Lab - Printing Resources and Datasources supported by any Terraform Provider 
```
cd ~/terraform-2428-feb2025
git pull
cd Day4/custom-terraform-provider/test-custom-terraform-docker-provider

terraform providers schema -json | jq
```

Expected output
![image](https://github.com/user-attachments/assets/6b5e7265-7c64-42eb-a686-3dd4bc44f47f)
![image](https://github.com/user-attachments/assets/5a00f5e0-b18a-4626-8ada-e328f3a042f4)

## Info - Terraform Provider Development best practices and recommended naming conventions
<pre>
- Provider name must be terraform-provider-nameoftheprovider, must be all lower case
- Resource names must start with nameoftheprovider-resource i.e docker_container, docker is the provider name while the resource managed is container
- resource and data source name must be all lowercase separated by underscore, and recommened to restrict to 2 or 3 words at the max
</pre>


## Info - Commons causes of Configuration Drift
<pre>
- manual changes
- inconsistent and manual deployments
- external dependencies
- difference in environments
- lack of version control
- poor or lack of documentation
</pre>

## Info - Risks associated with Configuration Drifts
<pre>
- Security Vulnerabilities
- Performance issues
- Compliance issues
- Makes troubleshooting very difficult
- increased down-time
- decreased reliability
- poor user-experience
</pre>

## Info - Solution to Configuration Drifts
<pre>
- Using Version Control
- Continuous Integration
- Use Configuration Management Tools to override manual changes in continuous fashion
- Use Infrastructure as Code Tools to override manual changes
- Regularly monitor and audit infrastructure
</pre>

## Info - Monitoring
<pre>
- is the process of collecting and analyzing data from IT infrastructure, system and processes
- using the data to improve business outcomes and drive value to the organization
- collects data to help keep your infrastructure up and running without any downtime
- Tools to
  - Data Collection (Logs)
  - Alerting
  - Remediation
  - Agent based monitoring
  - Agentless monitoring
  - Collecting Metrics
- Examples
  - Dynatrace
  - DataDog
  - Splunk
  - Prometheus & Grafana
</pre>

## Continuous Integration (CI) Overview
<pre>
- it is a fail-fast engineering process to find issues early 
- when bugs are detected early during development phase, they are easy to fix
- cost of fixing the bugs is also relatively cheaper
- it is similar SCRUM daily stand-up meeting, which is an inspect and adapt meeting
- the team that follows CI/CD, the engineers would be pushing code to version control several times a day
- code would be integrated many times a day
- Jenkins or similar CI/CD server will grab the latest code, they trigger a build, as part of the build, automated test cases would be executed to verify if the new code is as expected, if the new code is breaking any existing functionality.
- the build that was triggered due to new code integration succeeds, it means no functionality is broken, everything works as expected
- continous frequent feedback is given by CI/CD builds, eventually improving the code quaility and functional quality
- CI helps confidently deliver releases in a short amount of time
- Unit and Integration is the scope of CI
</pre>

## Continuous Deployment (CD) Overview
<pre>
- Once the dev release suceeds all the automated test cases added by dev team, it is automatically promoted for QA testing
- the dev certified release binaries are deployed automatically to QA environment for further automated QA testing
- the QA would automate, end to end functionality test, regression test, smoke test, performance test, stress test, component/API test, etc
</pre>

## Continuous Delivery (CD) Overview
<pre>
- the QA certified build(release) is automatically deployed into production or pre-prod environment
- in the pre-prod environment the customer or the Operations team would verify if the new release is working as expected
- especially fintech companies, after manual approval the binaries could go live in production environment
</pre>

## What is Jenkins?
- is a build automation server
- used mainly for CI/CD Builds
- this was developed in Java by Kohsuke Kawaguchi, former employee of Sun Microsystems
- Initially it was developed as Hudson is an opensource project
- Then Oracle acquired Sun Microsystems, then part of Hudson including Kohsuke Kawaguchi had quit Oracle
  created a new branch from Hudson called Jenkins
- The other part of the Hudson team they continue to develop the product as Hudson
- There is lot common code between Hudson and Jenkins
- More than 10000 active contributors are there for Jenkins
- Many other software vendors got inspired by Jenkins similar products came out in market like Bamboo, Team City, Microsoft TFS, etc.,
- Jenkins supports CI/CD build for products built in any software stack
  
## What is Cloudbees?
- Cloudbees is the enterprise paid variant of Jenkins
- Feature wise Jenkins and Cloudbees pretty much they are same
- We get support for Cloudbees while we only get community support for Jenkins
- Cloudbees is more stable as it is a paid version
  
## Jenkins Alternatives
- Bamboo
- Team City
- Cloudbees ( Enterprise Jenkins )
- Microsoft Team Foundation Server (TFS)

## Lab - Download Jenkins war file
Download the Generic Java Package (war) file from the left side (LTS)
<pre>
cd ~/Downloads
wget https://get.jenkins.io/war-stable/2.492.1/jenkins.war
</pre>

Expected output


## Lab - Launching Jenkins from terminal
```
cd ~/Downloads
java -jar ./jenkins.war
```

Expected output
![image](https://github.com/user-attachments/assets/a844181c-d7ba-4de7-bd79-8b6ada77f8ac)
![image](https://github.com/user-attachments/assets/1647eb41-a428-4b67-b44c-a7d263399a4e)
![image](https://github.com/user-attachments/assets/e990041c-9738-4cd6-9ac1-2ceadc022a1d)


## Lab - Accessing Jenkins Dashboard from your RPS Cloud machine chrome web browser
<pre>
http://localhost:8080  
</pre>

Expected output
![image](https://github.com/user-attachments/assets/9fb750bd-f98e-494b-9517-8a160ee68c99)

Copy the path from the above screen and open a separate tab in chrome browser to see the initial admin password
![image](https://github.com/user-attachments/assets/1bc36eea-7975-4aff-a258-1a06a001951c)
Copy this initial admin password and paste the same in Jenkins Tab
![image](https://github.com/user-attachments/assets/f3fbf663-c995-4e44-a8e5-fde54a770a67)
![image](https://github.com/user-attachments/assets/97f952b0-5d6a-41b3-8d66-1459e184bddd)
Click "Continue"
![image](https://github.com/user-attachments/assets/fbb7acad-5af3-41d3-bf39-0118f135ae60)
Click "Install suggested plugins"
![image](https://github.com/user-attachments/assets/4d02d877-0a5c-43a9-8e9c-406ca2ffc28c)
![image](https://github.com/user-attachments/assets/9b6b19b5-3aeb-476d-bdb2-7b7bfdaaef04)
![image](https://github.com/user-attachments/assets/19200eb4-7e0b-4cc6-a240-3a3ebe2b25ea)
![image](https://github.com/user-attachments/assets/f16b84a1-cfaf-4d7c-aacb-e70927eb3435)
You can type in your name and details requested
![image](https://github.com/user-attachments/assets/6824bedd-aa74-4293-ae64-3c4294658ef9)
![image](https://github.com/user-attachments/assets/42773f53-d435-44c4-9788-39d78656b7c6)
Click "Save and Continue"
![image](https://github.com/user-attachments/assets/6148df31-e7b7-465f-b64c-b804e7a4fa55)
Click "Save and Finish"
![image](https://github.com/user-attachments/assets/43424e6c-bc3e-42b1-9e94-626031a071d5)
Click "Start using Jenkins"
![image](https://github.com/user-attachments/assets/6cab67e6-9442-41ad-9fc5-14403c7f0406)

## Lab - Let's create a DevOps CI/CD Pipeline
<pre>
- We will be creating Declarative Jenkins DevOps Pipeline  
- The DevOps pipeline will have couple of stages
  - Clone source from GitHub (This will taken care by Jenkins out of the box)
  - Second stage - We will build our Custom Ansible Ubuntu Node Container Image
  - Third stage - We will build our custom docker terraform provider and install it
  - Fourth stage - We will run the terraform manifest script that provisions a container using the custom image we built in step 2
    - Once it is able to provision a container using custom image built in step 2
    - Terraform will invoke ansible playbook to perform configuration management i.e install nginx web server on it
  - We will get build report
</pre>

Let's navigate to Jenkins Dashboard in RPS Ubuntu cloud machine Chrome Web browser
![image](https://github.com/user-attachments/assets/3319dcc7-5355-4500-88e7-7ced89f5f11e)
In the middle of the screen, you can click "Create a job"
![image](https://github.com/user-attachments/assets/6df8d4ee-9171-4202-bc81-41c981b13d64)
Select the "Pipeline" and under "Enter an item name" type DevOpsCICDPipeline
![image](https://github.com/user-attachments/assets/f5d2d653-ae7c-487c-888a-501f8658457e)
Click "Ok"
![image](https://github.com/user-attachments/assets/500e24c7-b764-442a-a6f4-06fad39bcc3f)
General Section
![image](https://github.com/user-attachments/assets/87134d7c-e660-48e4-bc8e-bdf9317bb724)

Triggers Section
![image](https://github.com/user-attachments/assets/c462b3cb-48c8-424d-8c66-ba576da739a8)
Select "Poll SCM" checkbox and under schedule type "H/02 * * * *"
![image](https://github.com/user-attachments/assets/a6178c7e-4e5b-471a-97ea-50715c1e0dca)

Pipeline Section
![image](https://github.com/user-attachments/assets/59eb88af-a641-486f-b65c-9acc08f00902)
Click Definition
![image](https://github.com/user-attachments/assets/647d6c81-7364-4114-82a2-0936c24b598c)
Select "Pipeline Script from SCM"
![image](https://github.com/user-attachments/assets/e41594e3-d0c9-445b-a8f8-d0e5b6f3f3b1)
Under SCM, select "Git"
![image](https://github.com/user-attachments/assets/0d381c75-2d68-4073-8d62-684876a75c44)
![image](https://github.com/user-attachments/assets/b1337742-1710-4149-a8de-9211c00e801a)
Under Repository url type "https://github.com/tektutor/terraform-2428-feb2025.git"
![image](https://github.com/user-attachments/assets/6a1a93d9-6286-4546-a3a7-41336cd44248)
Under "Branch specifier" replace "master" with "main" branch
![image](https://github.com/user-attachments/assets/d78518eb-8cb0-4dbd-9385-3190f8473069)
Under Script Path, update "Day5/DevOpsCICDPipeline\Jenkinsfile"
![image](https://github.com/user-attachments/assets/cbeee504-c057-484e-a205-75da1fa5d1fd)
Click "Save"
Build Progress
![image](https://github.com/user-attachments/assets/26c04828-e23c-498a-8ca1-f1944065d1bc)
![image](https://github.com/user-attachments/assets/423b6452-9db3-432e-899c-759eaa766690)
![image](https://github.com/user-attachments/assets/2f914ca5-ad71-4fe9-bcd2-515ea991b2a9)
![image](https://github.com/user-attachments/assets/38d262ad-e934-4dbd-a9d7-5462c84cf30b)
![image](https://github.com/user-attachments/assets/b34b397c-2021-4b86-965c-8cd29a89ef18)

