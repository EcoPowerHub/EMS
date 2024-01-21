
This package provides a Golang implementation of an Equipment class. This package contains one struct named `Equipment` that represents an equipment connected to a Modbus TCP or RTU slave. It has two methods, `Configure()` and `AddOrRefreshData()`, both of which are used for configuring the equipment's connection to the slave and read its values respectively. The package also contains a struct named `readings` that holds the current state of the PV power generation in Watts.

The Equipment struct has the following fields:
- logger: A zerolog.Logger instance for logging purposes
- mc: A modbus.ModbusClient pointer to communicate with a Modbus slave
- state: An EquipmentState value that represents the current state of the equipment
- host: The address (IP or device name) of the Modbus TCP/RTU slave
- readings: A struct containing a float64 field for P_w
- lastRead: A time.Time variable to store the timestamp of the last successful reading

The Configure() method takes no arguments and returns an error if it fails to open a connection with the Modbus slave, otherwise sets the state to `common.EquipmentStateOnline`. It creates a new modbus client using `modbus.NewClient` and sets its URL to the value of the `host` field.

The AddOrRefreshData() method reads register 0 on the Modbus slave and stores the result in the `readings` struct, converting it to a float64. If an error occurs, it logs the error and sets the state to `common.EquipmentStateUnreachable`. Otherwise, it sets the state to `common.EquipmentStateOnline`.

The Read() method returns a map with a single key-value pair for the PV power generation in Watts (`io.PV`). The value of this field is obtained from the `readings` struct and the timestamp is set to the current Unix time in microseconds using `time.Now()`.

The Write() method takes a map containing a single key-value pair for PV power generation, but since it does not modify any registers on the slave, it returns an empty error. This method can be implemented if the equipment needs to write values to the slave in the future.
