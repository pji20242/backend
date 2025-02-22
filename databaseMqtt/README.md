

# to add connector: 

```sql
INSERT INTO mqttusers (username, pw, mosquitto_super) 
VALUES ('connector', 'PBKDF2$sha256$10000$ywt6X6oqPc3Zz6lw$kbv4Xl4HJbGKp1vtEvAC2T3+pmbTmhrq', true);
```

# to add device: 

```sql
INSERT INTO mqttusers (username, pw, mosquitto_super)
VALUES ('device', 'PBKDF2$sha256$10000$WeleOC9X8YQFpSO/$rQ2n4Zt6mdneOcSvFf0Nrnfv2iGqVKid', true);
```

## to verify it: 

```sql
SELECT * FROM mqttusers;
```

# to add ACL: 

```sql
INSERT INTO mqttacls (username, topic, rw) 
VALUES ('connector', 'pji3', 1);
```

```sql
INSERT INTO mqttacls (username, topic, rw) 
VALUES ('device', 'pji3', 1);
```

## to verify it: 

```sql
SELECT * FROM mqttacls;
```