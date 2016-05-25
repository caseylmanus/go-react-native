/**
 * Sample React Native App
 * https://github.com/facebook/react-native
 */

import React, {
  AppRegistry,
  Component,
  StyleSheet,
  Text,
  View,
  NativeModules
} from 'react-native';

class GoReact extends Component {
  constructor() {
    super();
    this.state = {};
    fetch("http://localhost:8080/").then((result) => {
      return result.text();
    }).then((text) => {
      this.setState({ httpMessage: text });
    });
    var ws = new WebSocket('ws://localhost:8080/websocket');

    ws.onopen = () => {
      // connection opened
      ws.send('something');
    };

    ws.onmessage = (e) => {
      this.setState({ webSocketMessage: e.data });
    };
    NativeModules.Go.sayHello((data) => {
      this.setState({nativeBridgeMessage: data});
    });
  }
  render() {
    return (
      <View style={styles.container}>
        <Text style={styles.welcome}>
          From HTTP: {this.state.httpMessage}
        </Text>
        <Text style={styles.welcome}>
          From WebSocket: {this.state.webSocketMessage}
        </Text>
         <Text style={styles.welcome}>
          From Native Bridge: {this.state.nativeBridgeMessage}
        </Text>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: '#F5FCFF',
  },
  welcome: {
    fontSize: 20,
    textAlign: 'center',
    margin: 10,
  },
  instructions: {
    textAlign: 'center',
    color: '#333333',
    marginBottom: 5,
  },
});

AppRegistry.registerComponent('GoReact', () => GoReact);
