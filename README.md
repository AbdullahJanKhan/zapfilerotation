# zapfilerotation

This package is for uber zap logger. It can be used as a file rotation package which enables users to rotate file based on time as well as size

It takes following input parameters

1. complete file path
2. time interval duration
3. max size of file in mega bytes

and returns file rotation instance

# Example

filename = "./logger.log"

writeFile := zapfilerotation.NewTimeRotationWriter(filename, time.Minute, 2)

core := zapcore.NewCore(
zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
zapcore.NewMultiWriteSyncer(zapcore.AddSync(writeFile)),
zap.InfoLevel,
)

logger := zap.New(core)

# Contribution
This project was born due to my professional needs but i would like to improve it, if anyone wants to join me in this journey just create a PR or raise an issue.