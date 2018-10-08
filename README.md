# WIP skill-sample-go-fact

This is an example fact skill implemented as an AWS Lambda in Go. It is not currently at parity with [skill-sample-nodejs-fact](https://github.cot/alexa/skill-sample-nodejs-fact), but should serve as a reasonable start.

## Quick[ish]start
 
This requires the use of [Serverless Framework](https://serverless.com/) to deploy your build. It's not mandatory, but you'll have to upload the build zip manually.

1. Set up the Alexa Skill in the console to get the Skill ID.
1. In the project root, copy `env.yml.template` to `env.yml` and provid the Skill ID for the `ALEXA_SKILL_ID` environment variable. 
1. Run `make` to build your project.
1. Run `sls deploy` to deploy your project.
1. Finish configuring your skill.

## Implementation Notes

With no official skills kit in Go, the request and response is handled manually as a learning exercise. However, here are some good options in no particular order to not reinvent the wheel:

* [alexa-sdk-go](https://gitlab.com/dasjott/alexa-sdk-go)
* [alexa-skills-kit-golang](https://github.com/ericdaugherty/alexa-skills-kit-golang)
* [alexa-go](https://github.com/arienmalec/alexa-go)

