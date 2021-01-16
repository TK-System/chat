/* talk edit */
let isMySelf = true;

let sendBtn = document.getElementById('sendBtn');
sendBtn.addEventListener('click', function () {

    let inputMessage = document.getElementById('inputMessage');
    let messageText = inputMessage.value.replace(/\r\n|\r/g, "\n");
    console.log("inputmessage:" + messageText);
    if (messageText == '') {
        return;
    }

    let messageBox = document.createElement('div');

    if (isMySelf) {
        messageBox.classList.add('box-right');
    } else {
        messageBox.classList.add('box-left');
    }

    //textarea1 no kaigyou dekitenai!!!!!
    var lines = messageText.split('\n');

    //height
    const messageBoxHeight1Line = 41;
    const messageBoxHeight2Line = 62; //padding(20px) + lines *(textsize(16) * 1.3)
    const upHeight = 21 //(textsize(16) * 1.3)

    var lengthArray = [];
    //width
    for (var i = 0; i < lines.length; i++) {
        console.log(charcount(lines[i]));
        lengthArray.push(charcount(lines[i]));
        console.log("test:" + lengthArray[i]);
    }

    //Max mozisuu
    var max = lengthArray.reduce(function (a, b) {
        return Math.max(a, b);
    });
    console.log("max:" + max);

    var message = document.createElement('textarea');
    //message.setAttribute("cols", "60");
    message.setAttribute("style", "overflow:hidden");
    message.setAttribute("disabled", "true");
    message.style.width = 30 + max * 9.8 + "px";
    console.log("width:" + message.style.width);
    if (lines.length === 1) {
        message.style.height = messageBoxHeight1Line + "px";
        console.log("height:" + message.style.height);
        console.log("kaigyou : nasi");
    } else {
        console.log("kaigyou: " + lines.length);
        message.style.height = messageBoxHeight2Line + (lines.length - 2) * upHeight + "px";
        console.log("height:" + message.style.height);
    }
    message.textContent = messageText;
    console.log("message2:" + message.textContent);
    message.classList.add('message-box');

    if (isMySelf) {
        message.classList.add('green');
        message.style.textAlign = "right";
    } else {
        message.classList.add('white');
    }

    messageBox.appendChild(message);

    let room = document.getElementById('room');

    room.appendChild(messageBox);

    isMySelf = !isMySelf;

    inputMessage.value = '';
    inputMessage.style.height = "62px";
})

/*
textarea height control
*/
function textAreaHeightSet(argObj) {
    console.log("height controll start!");
    argObj.style.height = "62px";
    var wScrollHeight = parseInt(argObj.scrollHeight);
    var wLineH = parseInt(argObj.style.lineHeight.replace(/px/, ''));
    if (wScrollHeight < (wLineH * 2)) {
        wScrollHeight = (wLineH * 2);
    }
    argObj.style.height = wScrollHeight + "px";
}

/*
byteCount
*/
var charcount = function (str) {
    len = 0;
    str = escape(str);
    for (i = 0; i < str.length; i++, len++) {
        if (str.charAt(i) == "%") {
            if (str.charAt(++i) == "u") {
                i += 3;
                len++;
            }
            i++;
        }
    }
    return len;
}