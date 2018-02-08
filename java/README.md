# Bosch IoT Things :: Java example clients

Find more documentation about the Things Java Client here: [Bosch IoT Things Wiki](https://things.s-apps.de1.bosch-iot-cloud.com/dokuwiki/doku.php?id=005_dev_guide:005_java_api:things-client_instantiation)

Bosch IoT Things provides an official Java client which is accessible on a public Maven repository. In order to use this
repository, add following information either to your Maven `settings.xml` or to your `pom.xml` files:

## Maven setup

```xml
<repositories>
    <repository>
       <id>bosch-si-public</id>
       <url>https://maven.bosch-si.com/content/repositories/bosch-releases</url>
       <releases>
          <enabled>true</enabled>
          <updatePolicy>never</updatePolicy>
       </releases>
       <snapshots>
          <enabled>true</enabled>
          <updatePolicy>daily</updatePolicy>
       </snapshots>
    </repository>
</repositories>
```

## Maven dependency

The Things client dependency is:

```xml
<dependency>
   <groupId>com.bosch.iot.things.client</groupId>
   <artifactId>things-client</artifactId>
   <version>3.0.0</version>
</dependency>
``` 

### OSGi

If you require an OSGi bundle, you can use instead:

```xml
<dependency>
   <groupId>com.bosch.iot.things.client</groupId>
   <artifactId>things-client-osgi</artifactId>
   <version>3.0.0</version>
</dependency>
``` 

The following `.zip` file contains all required OSGi dependencies which must also be present in the OSGi environment:
[things-client-osgi-3.0.0-bundles.zip](https://maven.bosch-si.com/content/repositories/bosch-releases/com/bosch/iot/things/client/things-client-osgi/3.0.0/things-client-osgi-3.0.0-bundles.zip)
