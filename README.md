# WIP skill-sample-go-fact

This is an example fact skill implemented as an AWS Lambda in Go. It is not currently at parity with [skill-sample-nodejs-fact](https://github.cot/alexa/skill-sample-nodejs-fact), but should serve as a reasonable start.

## Quick[ish]start
 
This requires the use of [Serverless Framework](https://serverless.com/) to deploy your build. It's not mandatory, but you'll have to upload the build zip manually.

1. Set up the Alexa Skill in the console to get the Skill ID.
1. In the project root, copy `env.yml.template` to `env.yml` and provid the Skill ID for the `ALEXA_SKILL_ID` environment variable. 
1. Run `make` to build your project.
1. Run `sls deploy` to deploy your project.
1. Finish configuring your skill.

## Invoking via CLI

### Via Serverless Framework
```sh
$ sls invoke -f skill
{
    "version": "1.0",
    "response": {
        "outputSpeech": {
            "type": "PlainText",
            "text": "Soccer is a British term appearing in the 1880s \n\t\tas an abbreviation for the word association, as in association football,  \n\t\ta formal name used to distinguish itself from rugby football."
        }
    }
}
```

### Via AWS CLI
```sh
$ aws lambda invoke --function-name <lambda name or ARN> --region <region> outfile.txt
{
    "StatusCode": 200
}
$ more outfile.txt
{"version":"1.0","response":{"outputSpeech":{"type":"PlainText","text":"Soccer i
s a British term appearing in the 1880s \n\t\tas an abbreviation for the word as
sociation, as in association football,  \n\t\ta formal name used to distinguish 
itself from rugby football."}}}
```

## Implementation Notes

With no official skills kit in Go, the request and response is handled manually as a learning exercise. However, here are some good options in no particular order to not reinvent the wheel:

* [alexa-sdk-go](https://gitlab.com/dasjott/alexa-sdk-go)
* [alexa-skills-kit-golang](https://github.com/ericdaugherty/alexa-skills-kit-golang)
* [alexa-go](https://github.com/arienmalec/alexa-go)

