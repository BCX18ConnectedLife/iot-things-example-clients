/*
 * Bosch SI Example Code License Version 1.0, January 2016
 *
 * Copyright 2018 Bosch Software Innovations GmbH ('Bosch SI'). All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without modification, are permitted provided that the
 * following conditions are met:
 *
 * 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following
 * disclaimer.
 *
 * 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the
 * following disclaimer in the documentation and/or other materials provided with the distribution.
 *
 * BOSCH SI PROVIDES THE PROGRAM 'AS IS' WITHOUT WARRANTY OF ANY KIND, EITHER EXPRESSED OR IMPLIED, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE. THE ENTIRE RISK AS TO THE
 * QUALITY AND PERFORMANCE OF THE PROGRAM IS WITH YOU. SHOULD THE PROGRAM PROVE DEFECTIVE, YOU ASSUME THE COST OF ALL
 * NECESSARY SERVICING, REPAIR OR CORRECTION. THIS SHALL NOT APPLY TO MATERIAL DEFECTS AND DEFECTS OF TITLE WHICH BOSCH
 * SI HAS FRAUDULENTLY CONCEALED. APART FROM THE CASES STIPULATED ABOVE, BOSCH SI SHALL BE LIABLE WITHOUT LIMITATION FOR
 * INTENT OR GROSS NEGLIGENCE, FOR INJURIES TO LIFE, BODY OR HEALTH AND ACCORDING TO THE PROVISIONS OF THE GERMAN
 * PRODUCT LIABILITY ACT (PRODUKTHAFTUNGSGESETZ). THE SCOPE OF A GUARANTEE GRANTED BY BOSCH SI SHALL REMAIN UNAFFECTED
 * BY LIMITATIONS OF LIABILITY. IN ALL OTHER CASES, LIABILITY OF BOSCH SI IS EXCLUDED. THESE LIMITATIONS OF LIABILITY
 * ALSO APPLY IN REGARD TO THE FAULT OF VICARIOUS AGENTS OF BOSCH SI AND THE PERSONAL LIABILITY OF BOSCH SI'S EMPLOYEES,
 * REPRESENTATIVES AND ORGANS.
 */

'use strict';

const WebSocket = require('ws');
const cli = require('./things_websocket_cli');
const readline = require('readline');

const rl = readline.createInterface({
                                        input: process.stdin,
                                        output: process.stdout
                                    });

const options = cli.options();
const endpointUrl = 'wss://things.s-apps.de1.bosch-iot-cloud.com/ws/2';
const wsOptions = {
    headers: {
        'Authorization': `Basic ${new Buffer(`${options.username}:${options.password}`).toString('base64')}`,
        'x-cr-api-token': options.apiToken
    }
};

log(`Connecting to ${endpointUrl}`);

const ws = new WebSocket(endpointUrl, wsOptions);

ws.on('open', () => {
    log('Connected');
    keepAlive(ws, 30000);

    rl.on('line', (input) => {
        if (JSON.parse(input)) {
            ws.send(input);
        } else {
            log('Input needs to be valid JSON');
        }
    });
    rl.on('close', () => ws.terminate());
});

ws.on('message', (data) => log(`Received: ${data}`));

ws.on('close', () => {
    log('Disconnected');
    process.exit(0);
});

function keepAlive(socket, heartbeat) {
    setInterval(function () {
        log('Sending keepalive');
        socket.send(new ArrayBuffer(0));
    }, heartbeat);
}

function log(message) {
    console.log(`${new Date().toISOString()} [WebSocket] ${message}`);
}
