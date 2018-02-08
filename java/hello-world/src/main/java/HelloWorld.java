/*
 *                                            Bosch SI Example Code License
 *                                              Version 1.0, January 2016
 *
 * Copyright 2018 Bosch Software Innovations GmbH ("Bosch SI"). All rights reserved.
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
 * BOSCH SI PROVIDES THE PROGRAM "AS IS" WITHOUT WARRANTY OF ANY KIND, EITHER EXPRESSED OR IMPLIED, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE. THE ENTIRE RISK AS TO
 * THE QUALITY AND PERFORMANCE OF THE PROGRAM IS WITH YOU. SHOULD THE PROGRAM PROVE DEFECTIVE, YOU ASSUME THE COST OF
 * ALL NECESSARY SERVICING, REPAIR OR CORRECTION. THIS SHALL NOT APPLY TO MATERIAL DEFECTS AND DEFECTS OF TITLE WHICH
 * BOSCH SI HAS FRAUDULENTLY CONCEALED. APART FROM THE CASES STIPULATED ABOVE, BOSCH SI SHALL BE LIABLE WITHOUT
 * LIMITATION FOR INTENT OR GROSS NEGLIGENCE, FOR INJURIES TO LIFE, BODY OR HEALTH AND ACCORDING TO THE PROVISIONS OF
 * THE GERMAN PRODUCT LIABILITY ACT (PRODUKTHAFTUNGSGESETZ). THE SCOPE OF A GUARANTEE GRANTED BY BOSCH SI SHALL REMAIN
 * UNAFFECTED BY LIMITATIONS OF LIABILITY. IN ALL OTHER CASES, LIABILITY OF BOSCH SI IS EXCLUDED. THESE LIMITATIONS OF
 * LIABILITY ALSO APPLY IN REGARD TO THE FAULT OF VICARIOUS AGENTS OF BOSCH SI AND THE PERSONAL LIABILITY OF BOSCH SI'S
 * EMPLOYEES, REPRESENTATIVES AND ORGANS.
 */

import java.util.concurrent.ExecutionException;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.TimeoutException;

import org.eclipse.ditto.model.base.json.JsonSchemaVersion;

import com.bosch.iot.things.client.ThingsClientFactory;
import com.bosch.iot.things.client.configuration.CommonConfiguration;
import com.bosch.iot.things.client.configuration.CredentialsAuthenticationConfiguration;
import com.bosch.iot.things.client.configuration.ProviderConfiguration;
import com.bosch.iot.things.client.messaging.MessagingProviders;
import com.bosch.iot.things.clientapi.ThingsClient;

/**
 * Shows a simple "Hello World" of Bosch IoT Things via the Java Client.
 */
public final class HelloWorld {

    private static final String USERNAME = "TODO insert your username here";
    private static final String PASSWORD = "TODO insert your password here";
    private static final String THINGS_API_TOKEN = "TODO insert your Things API-Token here";

    /**
     * Entry point.
     */
    public static void main(final String... args) throws InterruptedException, ExecutionException, TimeoutException {

        final ThingsClient client = initializeThingsClient();

        // First step: creating a Thing + updating a property afterwards
        client.twin()
                .create() // create Thing with unique ID
                .thenApply(createdThing -> {
                    System.out.println("Thing was created: " + createdThing);
                    return createdThing.getId().get();
                })
                .thenApply(thingId ->
                        client.twin()
                                .forId(thingId)
                                .putAttribute("foo", 42)
                )
                .thenAccept(_void ->
                        System.out.println("Attribute updated successfully"));

        // block until we receive events from back-end:
        client.twin().startConsumption().get(10, TimeUnit.SECONDS);

        // Second step: register for changes to all Things
        // make changes to the Thing via HTTP in order for that handler to be called:
        client.twin()
                .registerForThingChanges("GLOBAL_HANDLER", thingChange -> {
                    System.out.println("Change for Thing with ID:" + thingChange.getThingId());
                    System.out.println("Change action: " + thingChange.getAction());
                    System.out.println("Thing: " + thingChange.getThing().orElse(null));
                });
    }

    private static ThingsClient initializeThingsClient() {
        final CredentialsAuthenticationConfiguration credentialsAuthenticationConfiguration =
                CredentialsAuthenticationConfiguration
                        .newBuilder()
                        .username(USERNAME)
                        .password(PASSWORD)
                        .build();

        final ProviderConfiguration providerConfiguration = MessagingProviders
                .thingsWebsocketProviderBuilder()
                .authenticationConfiguration(credentialsAuthenticationConfiguration)
                .build();

        final CommonConfiguration configuration = ThingsClientFactory.configurationBuilder()
                .apiToken(THINGS_API_TOKEN)
                .providerConfiguration(providerConfiguration)
                .schemaVersion(
                        JsonSchemaVersion.V_2) // explicitly define "V_2" in order to create Things in API 2 with Policies
                .build();

        return ThingsClientFactory.newInstance(configuration);
    }
}
