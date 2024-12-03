"use strict";
exports.__esModule = true;
var react_1 = require("react");
var react_native_1 = require("react-native");
var HomeScreen = function (_a) {
    var navigation = _a.navigation;
    var _b = react_1.useState(''), ipAddress = _b[0], setIpAddress = _b[1];
    return (react_1["default"].createElement(react_native_1.View, { style: styles.container },
        react_1["default"].createElement(react_native_1.Text, null, "Enter the IP address to ping:"),
        react_1["default"].createElement(react_native_1.TextInput, { style: styles.input, placeholder: "Enter IP address", value: ipAddress, onChangeText: setIpAddress, keyboardType: "numeric" }),
        react_1["default"].createElement(react_native_1.Button, { title: "Connect", onPress: function () { return navigation.navigate('Login', { ip: ipAddress }); } })));
};
exports["default"] = HomeScreen;
var styles = react_native_1.StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff',
        alignItems: 'center',
        justifyContent: 'center',
        padding: 16
    },
    input: {
        height: 40,
        borderColor: '#ccc',
        borderWidth: 1,
        borderRadius: 5,
        width: '100%',
        marginVertical: 10,
        paddingHorizontal: 10
    },
    responseContainer: {
        marginTop: 20,
        padding: 10,
        backgroundColor: '#f0f0f0',
        borderRadius: 5,
        width: '100%'
    },
    responseText: {
        fontWeight: 'bold',
        marginBottom: 5
    }
});
