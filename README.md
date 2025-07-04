start with `./scripts/start.sh`

faktory dashboard will be hosted at http://localhost:7420/ shortly after start

e.g. logs
```
╰─❯ ./scripts/start.sh
Faktory already started
Running Go app...
2025/07/04 02:37:37.418172 faktory_worker_go 1.7.0 PID 49870 now ready to process jobs
2025/07/04 02:37:37.418235 Using Faktory Client API 1.9.2
2025/07/04 10:37:37 [simulateMockJobs] pushing job. jobID=ecadf168-0d39-4971-9429-c5f3f266b337
2025/07/04 10:37:37 [simulateMockJobs] pushing job. jobID=51ce84f0-85b1-47de-8e21-ea09760464df
2025/07/04 10:37:37 [simulateMockJobs] pushing job. jobID=30276402-31b2-4849-928f-771708c3bcc2
2025/07/04 10:37:37 [SendOnboardingEmail] Working on job 30276402-31b2-4849-928f-771708c3bcc2
2025/07/04 10:37:37 [SendOnboardingEmail] Working on job ecadf168-0d39-4971-9429-c5f3f266b337
2025/07/04 10:37:38 [SendOnboardingEmail] Working on job 51ce84f0-85b1-47de-8e21-ea09760464df
2025/07/04 10:37:38 [simulateMockJobs] pushing job. jobID=8c8c192d-37ed-4f4d-831b-172a1f00d0cf
2025/07/04 10:37:38 [simulateMockJobs] pushing job. jobID=474118d3-c64f-4a41-9f9a-eb465f5cf0c5
2025/07/04 10:37:38 [simulateMockJobs] pushing job. jobID=c69df29b-19e9-46ff-b782-5744e92dd4c0
2025/07/04 10:37:38 [SendOnboardingEmail] Working on job c69df29b-19e9-46ff-b782-5744e92dd4c0
2025/07/04 02:37:38.634741 Error running send-onboarding-email job c69df29b-19e9-46ff-b782-5744e92dd4c0: helllo
2025/07/04 10:37:38 [simulateMockJobs] pushing job. jobID=1ac5e766-0100-4b6e-a5e8-80ab7e8ab9c3
2025/07/04 10:37:39 [SendOnboardingEmail] Working on job 8c8c192d-37ed-4f4d-831b-172a1f00d0cf
^C
2025/07/04 02:37:39.078807 Shutting down...
2025/07/04 10:37:39 [SendOnboardingEmail] Working on job 1ac5e766-0100-4b6e-a5e8-80ab7e8ab9c3
2025/07/04 10:37:39 [simulateMockJobs] pushing job. jobID=79571e58-1f10-429c-b388-477ab6dc2d66
2025/07/04 10:37:39 [SendOnboardingEmail] Working on job 474118d3-c64f-4a41-9f9a-eb465f5cf0c5
2025/07/04 10:37:39 [SendOnboardingEmail] Working on job 79571e58-1f10-429c-b388-477ab6dc2d66
2025/07/04 02:37:41.100260 Goodbye
```
