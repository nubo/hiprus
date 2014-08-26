hiprus
==========

Package hiprus provides a Hipchat hook for the [logrus] loggin package.

## Example

```Go
logrus.AddHook(&hiprus.HiprusHook{
	AuthToken:      "0e73972805c491b08bafef0d62704a",
	AcceptedLevels: hiprus.LevelThreshold(logrus.WarnLevel),
	RoomName:       "DevOps",
	Username:       "ExampleApp",
})
```

[logrus]: https://github.com/sirupsen/logrus
