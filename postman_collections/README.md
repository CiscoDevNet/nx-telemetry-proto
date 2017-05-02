POSTMAN chrome app reference:
=============================
 - https://www.getpostman.com/docs/introduction

Telemetry.postman_collection file:
==================================
    - This is a json file for the telemetry feature.

    - It has REST payloads for following.

          : Authentication.

          : Enable/Disable feature telemetry.

          : POST/GET/DELETE for certificate, sensor group, sensor path, destination group, destination profile, subscription.

Using postman collections:
==========================
    - Its a easyway to execute the REST operations on the management box.

    - User needs to import collection file, setup the global variable {{ip}} to the management ip of the box and start executing operations immediately.

    - Please refer https://github.com/postmanlabs/postman-app-support/wiki/Collections and import the collections.

    - Refer https://www.getpostman.com/docs/environments to see how to set global variable.