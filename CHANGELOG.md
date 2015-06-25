# Change Log
All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning](http://semver.org/).

## [1.1.0] - 2015-06-11
- Added validation for the complete querystring (for correct delimiter use)
- The psql output now only includes the wanted result value (psql -At) which is used for comparison
- Result values for matched and unmatched queries are now part of the check output
- Calling the check without any parameters will show the usage instead of "UNKNOWN - No parameters given!"

## [1.0.0] - 2015-06-10
- Initial release