"use strict";
exports.__esModule = true;
var react_1 = require("react");
var react_native_1 = require("react-native");
var vector_icons_1 = require("@expo/vector-icons");
var BottomNavBar = function (_a) {
    var navigation = _a.navigation;
    return (react_1["default"].createElement(react_native_1.View, { style: styles.navbarContainer },
        react_1["default"].createElement(react_native_1.TouchableOpacity, { onPress: function () { return navigation.navigate('Home'); }, style: styles.navButton },
            react_1["default"].createElement(vector_icons_1.Ionicons, { name: "home-outline", size: 24, color: "black" })),
        react_1["default"].createElement(react_native_1.TouchableOpacity, { onPress: function () { return navigation.navigate('Add'); }, style: styles.navButton },
            react_1["default"].createElement(vector_icons_1.Ionicons, { name: "add-circle-outline", size: 24, color: "black" })),
        react_1["default"].createElement(react_native_1.TouchableOpacity, { onPress: function () { return navigation.navigate('History'); }, style: styles.navButton },
            react_1["default"].createElement(vector_icons_1.Ionicons, { name: "time-outline", size: 24, color: "black" })),
        react_1["default"].createElement(react_native_1.TouchableOpacity, { onPress: function () { return navigation.navigate('Profile'); }, style: styles.navButton },
            react_1["default"].createElement(vector_icons_1.Ionicons, { name: "person-outline", size: 24, color: "black" }))));
};
var AreasScreen = function (_a) {
    var _b;
    var navigation = _a.navigation, route = _a.route;
    var areas = [
        { text: 'Upload every day', color: '#FF4D4D', icons: ['logo-github', 'time-outline'] },
        { text: 'Start Music!', color: '#4CAF50', icons: ['cloud-outline', 'logo-spotify'] },
        { text: 'Upload every day', color: '#9C27B0', icons: ['cloud-upload-outline', 'time-outline'] },
        { text: 'Stock photo!', color: '#2196F3', icons: ['mail-outline', 'logo-dropbox'] },
    ];
    var ip = (((_b = route.params) === null || _b === void 0 ? void 0 : _b.ip) || 'localhost').ip;
    return (react_1["default"].createElement(react_native_1.View, { style: styles.container },
        react_1["default"].createElement(react_native_1.Text, { style: styles.header }, "My AREAs"),
        react_1["default"].createElement(react_native_1.View, { style: styles.areasContainer }, areas.map(function (area, index) { return (react_1["default"].createElement(react_native_1.View, { key: index, style: [styles.areaBox, { backgroundColor: area.color }] },
            react_1["default"].createElement(react_native_1.Text, { style: styles.areaText }, area.text),
            react_1["default"].createElement(react_native_1.View, { style: styles.iconsContainer }, area.icons.map(function (icon, idx) { return (react_1["default"].createElement(vector_icons_1.Ionicons, { key: idx, name: icon, size: 24, color: "white", style: styles.areaIcon })); })))); })),
        react_1["default"].createElement(BottomNavBar, { navigation: navigation })));
};
var styles = react_native_1.StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: 'white',
        padding: 16
    },
    header: {
        fontSize: 24,
        fontWeight: 'bold',
        marginBottom: 16
    },
    areasContainer: {
        flexDirection: 'row',
        flexWrap: 'wrap',
        justifyContent: 'space-around'
    },
    areaBox: {
        width: 150,
        height: 150,
        borderRadius: 16,
        justifyContent: 'center',
        alignItems: 'center',
        marginBottom: 16
    },
    areaText: {
        color: 'white',
        fontSize: 16,
        fontWeight: 'bold',
        textAlign: 'center',
        marginBottom: 8
    },
    iconsContainer: {
        flexDirection: 'row',
        justifyContent: 'space-around',
        width: '80%'
    },
    areaIcon: {
        marginHorizontal: 4
    },
    navbarContainer: {
        flexDirection: 'row',
        justifyContent: 'space-around',
        alignItems: 'center',
        paddingVertical: 8,
        backgroundColor: '#f0f0f0',
        borderTopWidth: 1,
        borderTopColor: '#d0d0d0'
    },
    navButton: {
        alignItems: 'center',
        justifyContent: 'center'
    }
});
exports["default"] = AreasScreen;
