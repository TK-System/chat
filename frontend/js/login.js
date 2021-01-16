/*
get userInfo 
*/
document.getElementById("submit").onclick = function(){
    // ウェブストレージに対応している
    if (window.sessionStorage) {
        // ------------------------------------------------------------
        // セッションストレージオブジェクトを取得する
        // ------------------------------------------------------------
        var session_storage = window.sessionStorage;
        // ------------------------------------------------------------
        var str = document.getElementById('userid').value;
        window.sessionStorage.setItem("userId", str);
        // 文字列を読み込み
        var test = window.sessionStorage["userId"];
        // 出力テスト
        console.log(test);
    }
}