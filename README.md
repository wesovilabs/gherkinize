# Gherkinize

A command line tool that will help us to write better Gherkin scenarios.


## Purpose

The purpose of this project is mainly helping people to write **Gherkin scenarios** correctly.


## Getting Started


### Installing with Brew

```
brew tap wesovilabs/gherkinize
brew install gherkinize
```

Once the software is installed we can make use of it.

**gherkinize --help / gherkinize -h**

It displays the menu.

![alt text](https://github.com/wesovilabs/gherkinize/blob/master/doc/helpmenu.png "Gherkinize help")

**gherkinize --version / gherkinize -v**

It displays the current software version.

![alt text](https://github.com/wesovilabs/gherkinize/blob/master/doc/versionmenu.png "Gherkinize version")

**gherkinize --input [scenarios_dir] --config [gherkin-rules.toml] validate**

It validates our scenarios. Both arguments **input** and **config** are required.

- input: it must point to a directory that contains the scenarios to be validated. (So far it only supports directories path)
- config:  it is a **toml** configuration file path that must look like below. You can download [this](https://github.com/wesovilabs/gherkinize/blob/master/config/gherkin-rules.toml) and
modify the values at your convenience.

### Running as a Docker container

In case of you prefer run **gherkinize** with a Docker container you can do it by running the below command
```
    docker run \
    -v <PATH_TO_SCENARIOS_DIR>:/var/gherkinize/scenarios \
    -v <PATH_TO_DIRECTORY_WITH_RULES_CONFIGURATION_FILE_TOML>:/var/gherkinize/config \
    -it wesovilabs/gherkinize \
    /app/dist/gherkinize -i /var/gherkinize/scenarios  -c /var/gherkinize/config/<RULES_CONFIGURATION_FILE_TOML> validate
```

For example:
```
    docker run \
    -v /Users/Ivan/Sandbox/WesoviLabs/GO_WORKSPACE/src/github.com/wesovilabs/gherkinize/testdata/scenarios/:/var/gherkinize/scenarios \
    -v /Users/Ivan/Sandbox/WesoviLabs/GO_WORKSPACE/src/github.com/wesovilabs/gherkinize/config/:/var/gherkinize/config \
    -it wesovilabs/gherkinize \
    /app/dist/gherkinize -i /var/gherkinize/scenarios  -c /var/gherkinize/config/gherkin-rules.toml validate
```

## Gherkin rules configuration file

![alt text](https://github.com/wesovilabs/gherkinize/blob/master/doc/config.png "Gherkin Rules Validator")

**max_steps_per_scenario**: It determines the maximum number of steps per scenario. Many steps could make our scenarios not easy-to-read

**max_len_step**: Max len for steps. Steps with long text are usually hard to understand.

**empty_feature**: Set false if you don't want to allow empty features without scenarios.

**empty_scenario**: Set false if you don't want to allow empty scenarios without steps.

**strict**: This checks the below conditions:

* The first step must be GIVEN

* The keyword GIVEN can be only used once.

* WHEN can only be used after GIVEN statements

* THEN statement can only be used after WHEN statements


## Samples

Assuming you make use of default configuration file and have a directory with somce sceanarios.

For example, if you checkout this repository and run the below command:

    gherkinize -i testdata/scenarios/ -c config/gherkin-rules.toml  validate

the output will be the following:

### scenarioA.feature
![alt text](https://github.com/wesovilabs/gherkinize/blob/master/doc/scenarioa.png "Scenario A")

The output will be:

![alt text](https://github.com/wesovilabs/gherkinize/blob/master/doc/scenarioA_result.png "Scenario A Result")

### scenarioB.feature
![alt text](https://github.com/wesovilabs/gherkinize/blob/master/doc/scenarioB.png "Scenario B")

The output will be:

![alt text](https://github.com/wesovilabs/gherkinize/blob/master/doc/scenarioB_result.png "Scenario B Result")

### scenarioC.feature
![alt text](https://github.com/wesovilabs/gherkinize/blob/master/doc/scenarioC.png "Scenario B")

The output will be:

![alt text](https://github.com/wesovilabs/gherkinize/blob/master/doc/scenarioC_result.png "Scenario B Result")

### scenarioD.feature
![alt text](https://github.com/wesovilabs/gherkinize/blob/master/doc/scenarioD.png "Scenario B")

The output will be:

![alt text](https://github.com/wesovilabs/gherkinize/blob/master/doc/scenarioD_result.png "Scenario B Result")

### scenarioE.feature
![alt text](https://github.com/wesovilabs/gherkinize/blob/master/doc/scenarioE.png "Scenario B")

The output will be:

![alt text](https://github.com/wesovilabs/gherkinize/blob/master/doc/scenarioE_result.png "Scenario B Result")

### scenarioF.feature
![alt text](https://github.com/wesovilabs/gherkinize/blob/master/doc/scenarioF.png "Scenario B")

The output will be:

![alt text](https://github.com/wesovilabs/gherkinize/blob/master/doc/scenarioF_result.png "Scenario B Result")

## scenarioG.feature
![alt text](https://github.com/wesovilabs/gherkinize/blob/master/doc/scenarioG.png "Scenario B")

The output will be:

![alt text](https://github.com/wesovilabs/gherkinize/blob/master/doc/scenarioG_result.png "Scenario B Result")

## The code

This project is completely developed in **Go**. Please don't hate me so much if the code contains bad practices since I am just learning about Go coding **;-)**

### Understanding the project

If you wanna run the project locally I invite you to **[fork this repository](https://github.com/wesovilabs/gherkinize)** and check it out.

#### Makefile

You will find a **Makefile** in the project root.

* **make clean**: Remove compiled files.

* **make test**: Run the tests  (there's no tests yet... so Sorry!!!)

* **make gherkinize**: Build the runnable file.

* **make install**: Install locally.


## Built With

* [Go](https://golang.org/) - The go programming language

## Contributing

Please read [CONTRIBUTING.md](https://github.com/wesovilabs/gherkinize/blob/master/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/wesovilabs/gherkinize/tags).

## Authors

* **Iván Corrales Solera** - [Wesovilabs](http://www.wesovilabs.com)

See also the list of [contributors](https://github.com/wesovilabs/gherkinize/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
