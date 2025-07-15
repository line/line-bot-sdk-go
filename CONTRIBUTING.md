# How to contribute to LINE Bot SDK for Go project

First of all, thank you so much for taking your time to contribute! LINE Bot SDK for Go is not very different from any other open
source projects you are aware of. It will be amazing if you could help us by doing any of the following:

- File an issue in [the issue tracker](https://github.com/line/line-bot-sdk-go/issues) to report bugs and propose new features and
  improvements.
- Ask a question using [the issue tracker](https://github.com/line/line-bot-sdk-go/issues) (__Please ask only about this SDK__).
- Contribute your work by sending [a pull request](https://github.com/line/line-bot-sdk-go/pulls).

## Development

### Understand the project structure

The project structure is as follows:

- `linebot`: The main library code, organized into sub-packages by API functionality.
- `examples`: Example projects that demonstrate how to use the library.
- `generator`: OpenAPI-based code generation tools and templates.
- `script`: Utility scripts for development.

### Edit OpenAPI templates

Almost all code is generated with OpenAPI Generator based on [line-openapi](https://github.com/line/line-openapi)'s YAML files.
Thus, you cannot edit most code under the `linebot` directory directly.

You need to edit the custom generator templates under the `generator/src/main/resources` directory instead.

After editing the templates, run the `generate-code.py` script to generate the code, and then commit all affected files.
If not, CI status will fail.

When you update code, be sure to check consistencies between generated code and your changes.

### Add unit tests

We use Go's built-in testing framework. To run all tests with race detection and coverage reporting: `bash script/codecov.sh`

Especially for bug fixes, please follow this flow for testing and development:
1. Write a test before making changes to the library and confirm that the test fails.
2. Modify the code of the library.
3. Run the test again and confirm that it passes thanks to your changes.

### Run your code in your local

You can use the [example projects](examples) to test your changes locally before submitting a pull request.

## Contributor license agreement

When you are sending a pull request and it's a non-trivial change beyond fixing typos, please make sure to sign
[the ICLA (individual contributor license agreement)](https://cla-assistant.io/line/line-bot-sdk-go). Please
[contact us](mailto:dl_oss_dev@linecorp.com) if you need the CCLA (corporate contributor license agreement).