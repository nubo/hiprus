hiprus
==========

Package hiprus provides a Hipchat hook for the [logrus] loggin package.

Please see the [documentation].

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
[documentation]: http://godoc.org/github.com/nubo/hiprus

---
Version 2.0.0
