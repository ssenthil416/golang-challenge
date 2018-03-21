# golang-challenge

Skeleton for the golang challenge

# setup

Before you start, begin downloading the data files from:

https://drive.google.com/drive/folders/1yJQrNke6fiWeyeZR-IohbGIQSn72E7y8?usp=sharing

Clone this repo, you should get a ridiculously basic stub for the challenge. The dataservice package contains interfaces and struct that represent the data in the downloaded gob files.

Goals:
- ./run_test.sh outputs OK.
  - You'll have to modify TestMain, but that's the only one allowed.
- Implement a restful API that allows you to GET and DELETE entries from the DumbDictionary.
  - Response format and errors are up to you.
  - Use good standards and common sense.
- All requests to your api must be under 100ms.
  - Consider what would happen with high traffic or even bigger dumb dictionary.
  - Remember, there are no requirements for data persistance.
