"use client";

export default function Home() {
  // useEffect(() => {
  //   const message = {
  //     type: "initSearch",
  //     payload: null,
  //   };

  //   if (window.ReactNativeWebView) {
  //     window.ReactNativeWebView.postMessage(JSON.stringify(message));
  //   } else if (window.FlutterChannel) {
  //     window.FlutterChannel.postMessage(JSON.stringify(message));
  //   } else {
  //     console.warn("No WebView message handler found");
  //   }
  // }, []);

  const handleButtonClick = () => {
    const message = {
      type: "testMessage",
      payload: null,
    };

    if (window.ReactNativeWebView) {
      window.ReactNativeWebView.postMessage(JSON.stringify(message));
    } else if (window.FlutterChannel) {
      window.FlutterChannel.postMessage(JSON.stringify(message));
    } else {
      alert("Flutterとの接続はありません。");
    }
  };

  return (
    <div style={styles.container}>
      <div style={styles.card}>
        <h1 style={styles.title}>Initializing Search...</h1>
        <p style={styles.text}>アプリに初期化フラグを送信します・・・</p>
        <button onClick={handleButtonClick} style={styles.button}>
          値を送信する
        </button>
      </div>
    </div>
  );
}

const styles = {
  container: {
    height: "100vh",
    backgroundColor: "#f8fafc",
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
  },
  card: {
    backgroundColor: "#ffffff",
    padding: "2rem 3rem",
    borderRadius: "16px",
    boxShadow: "0 8px 24px rgba(0,0,0,0.08)",
    textAlign: "center",
    maxWidth: "400px",
    width: "90%",
  },
  title: {
    margin: "0 0 1rem",
    fontSize: "1.5rem",
    fontWeight: 600,
    color: "#0f172a",
  },
  text: {
    marginBottom: "0.5rem",
    color: "#334155",
  },
  subtext: {
    fontSize: "0.875rem",
    color: "#64748b",
  },
  button: {
    backgroundColor: "#2563eb",
    color: "#ffffff",
    padding: "0.75rem 1.5rem",
    fontSize: "1rem",
    fontWeight: 600,
    border: "none",
    borderRadius: "0.5rem",
    cursor: "pointer",
    transition: "background-color 0.3s",
  },
};
