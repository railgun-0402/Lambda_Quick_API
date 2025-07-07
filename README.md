# Lambda_Quick_API
## 概要
### 内容 
- LambdaとAPI Gatewayで動くAPIをGo + Echoで試しに作りたいだけ
- ReactのJotaiライブラリを使ってみたいので、Next.jsでUIを作成しAPIを実行できるようにする
- デプロイ方法は検討中

## 使用技術(作成中なので予定も含む)
### バックエンド
<img src="https://skillicons.dev/icons?i=go" /> <br /><br />

### その他インフラ
<img src="https://skillicons.dev/icons?i=docker,mysql,aws" /> <br /><br />

## tmp flutter code 

```
import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:webview_flutter/webview_flutter.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  final String webUrl = 'http://localhost:3000/'; // ← ここにNext.jsのURL

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter WebView',
      home: Scaffold(
        appBar: AppBar(title: const Text('WebView Sample')),
        body: WebViewContainer(url: webUrl),
      ),
    );
  }
}

class WebViewContainer extends StatefulWidget {
  final String url;

  const WebViewContainer({required this.url});

  @override
  State<WebViewContainer> createState() => _WebViewContainerState();
}

class _WebViewContainerState extends State<WebViewContainer> {
  late final WebViewController _controller;

  @override
  void initState() {
    super.initState();
    _controller = WebViewController()
      ..setJavaScriptMode(JavaScriptMode.unrestricted)
      ..addJavaScriptChannel(
        'FlutterChannel',
        onMessageReceived: (JavaScriptMessage message) {
          try {
            final data = json.decode(message.message);
            String jsonData = jsonEncode(data);

            if (data['type'] == 'initSearch') {
              debugPrint('✅ Flutterが initSearch を受信しました！');
              // 必要ならここで画面遷移やロジックを呼ぶ
            }
          } catch (e) {
            debugPrint('⚠️ メッセージパースエラー: $e');
          }
        },
      )
      ..loadRequest(Uri.parse(widget.url));
  }

  @override
  Widget build(BuildContext context) {
    return WebViewWidget(controller: _controller);
  }
}

```
