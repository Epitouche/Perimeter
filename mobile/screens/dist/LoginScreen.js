"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (_) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
exports.__esModule = true;
var react_1 = require("react");
var react_native_1 = require("react-native");
var LoginScreen = function (_a) {
    var _b;
    var navigation = _a.navigation, route = _a.route;
    var _c = react_1.useState(''), username = _c[0], setUsername = _c[1];
    var _d = react_1.useState(''), password = _d[0], setPassword = _d[1];
    var _e = react_1.useState({ username: '', password: '' }), errors = _e[0], setErrors = _e[1];
    var ip = ((_b = route.params) === null || _b === void 0 ? void 0 : _b.ip) || 'localhost';
    var handleLogin = function () { return __awaiter(void 0, void 0, void 0, function () {
        var hasError, newErrors, response, data, error_1;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    hasError = false;
                    newErrors = { username: '', password: '' };
                    if (!username) {
                        newErrors.username = 'Username is required';
                        hasError = true;
                    }
                    if (!password) {
                        newErrors.password = 'Password is required';
                        hasError = true;
                    }
                    setErrors(newErrors);
                    if (!!hasError) return [3 /*break*/, 7];
                    _a.label = 1;
                case 1:
                    _a.trys.push([1, 6, , 7]);
                    return [4 /*yield*/, fetch("http://" + ip + ":8080/api/v1/auth/login", {
                            method: 'POST',
                            headers: {
                                'Content-Type': 'application/json'
                            },
                            body: JSON.stringify({ username: username, password: password })
                        })];
                case 2:
                    response = _a.sent();
                    if (!response.ok) return [3 /*break*/, 4];
                    return [4 /*yield*/, response.json()];
                case 3:
                    data = _a.sent();
                    console.log('Data:', data);
                    navigation.navigate('AreaView', { ip: ip });
                    return [3 /*break*/, 5];
                case 4:
                    console.error('Error:', response.status);
                    _a.label = 5;
                case 5: return [3 /*break*/, 7];
                case 6:
                    error_1 = _a.sent();
                    console.error('Error:', error_1);
                    return [3 /*break*/, 7];
                case 7: return [2 /*return*/];
            }
        });
    }); };
    var switchToSignup = function () {
        console.log('Switch to signup');
        navigation.navigate('SignUp', { ip: ip });
    };
    return (react_1["default"].createElement(react_native_1.View, { style: styles.container },
        react_1["default"].createElement(react_native_1.Text, { style: styles.header }, "Log in"),
        react_1["default"].createElement(react_native_1.TextInput, { style: styles.input, placeholder: "Enter username", placeholderTextColor: "#aaa", value: username, onChangeText: function (text) { return setUsername(text); } }),
        errors.username ? (react_1["default"].createElement(react_native_1.Text, { style: styles.errorText }, errors.username)) : null,
        react_1["default"].createElement(react_native_1.TextInput, { style: styles.input, placeholder: "Enter password", placeholderTextColor: "#aaa", secureTextEntry: true, value: password, onChangeText: function (text) { return setPassword(text); } }),
        errors.password ? (react_1["default"].createElement(react_native_1.Text, { style: styles.errorText }, errors.password)) : null,
        react_1["default"].createElement(react_native_1.TouchableOpacity, null,
            react_1["default"].createElement(react_native_1.Text, { style: styles.forgotPassword }, "Forgot password?")),
        react_1["default"].createElement(react_native_1.TouchableOpacity, { style: styles.loginButton, onPress: handleLogin },
            react_1["default"].createElement(react_native_1.Text, { style: styles.loginButtonText }, "Log in")),
        react_1["default"].createElement(react_native_1.View, { style: styles.signUpContainer },
            react_1["default"].createElement(react_native_1.Text, { style: styles.newText }, "New?"),
            react_1["default"].createElement(react_native_1.TouchableOpacity, { onPress: switchToSignup },
                react_1["default"].createElement(react_native_1.Text, { style: styles.signUpText }, "Sign Up"))),
        react_1["default"].createElement(react_native_1.View, { style: styles.dividerContainer },
            react_1["default"].createElement(react_native_1.View, { style: styles.divider }),
            react_1["default"].createElement(react_native_1.Text, { style: styles.orText }, "or log in with"),
            react_1["default"].createElement(react_native_1.View, { style: styles.divider })),
        react_1["default"].createElement(react_native_1.View, { style: styles.socialIconsContainer },
            react_1["default"].createElement(react_native_1.Image, { source: { uri: 'https://img.icons8.com/color/48/google-logo.png' }, style: styles.socialIcon }),
            react_1["default"].createElement(react_native_1.Image, { source: { uri: 'https://img.icons8.com/ios-glyphs/50/github.png' }, style: styles.socialIcon }),
            react_1["default"].createElement(react_native_1.Image, { source: { uri: 'https://img.icons8.com/color/48/facebook.png' }, style: styles.socialIcon }))));
};
var styles = react_native_1.StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        paddingHorizontal: 20,
        backgroundColor: '#f9f9f9'
    },
    header: {
        fontSize: 32,
        fontWeight: 'bold',
        marginBottom: 20
    },
    input: {
        width: '100%',
        padding: 12,
        borderRadius: 20,
        borderWidth: 1,
        borderColor: '#ccc',
        marginBottom: 10,
        backgroundColor: '#fff'
    },
    forgotPassword: {
        alignSelf: 'flex-end',
        color: '#007BFF',
        marginBottom: 20
    },
    loginButton: {
        width: '100%',
        backgroundColor: '#000',
        padding: 12,
        borderRadius: 20,
        alignItems: 'center',
        marginBottom: 20
    },
    loginButtonText: {
        color: '#fff',
        fontWeight: 'bold'
    },
    signUpContainer: {
        flexDirection: 'row',
        alignItems: 'center',
        marginBottom: 20
    },
    newText: {
        marginRight: 5,
        color: '#555'
    },
    signUpText: {
        color: '#007BFF',
        fontWeight: 'bold'
    },
    dividerContainer: {
        flexDirection: 'row',
        alignItems: 'center',
        marginVertical: 20,
        width: '100%'
    },
    divider: {
        flex: 1,
        height: 1,
        backgroundColor: '#ccc'
    },
    orText: {
        marginHorizontal: 10,
        color: '#555'
    },
    socialIconsContainer: {
        flexDirection: 'row',
        justifyContent: 'center',
        marginTop: 10
    },
    socialIcon: {
        width: 40,
        height: 40,
        marginHorizontal: 10
    },
    errorText: {
        color: 'red',
        fontSize: 12,
        alignSelf: 'flex-start',
        marginBottom: 10
    }
});
exports["default"] = LoginScreen;
