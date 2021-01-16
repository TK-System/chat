      // userId from session
      //var userId = window.sessionStorage["userId"];
		  // 出力テスト
      //console.log(userId);
      
      //GetFriendAPI
      //const constFrinedListURL = '/friend/' + userId;
      //var friendList;
      //fetch(constFrinedListURL.then((res)=>{
        //return (res.json());
      //})
      //.then((json)=>{
        //syori
        //friendList = JSON.parse(json);
        //console.log(friendList);
      //})

		  
     
      window.onload = function () {

        // HTMLFormElement オブジェクト
        var form = document.getElementById("form");
        // 出力テスト
        console.log(form);
  
        //divTag for accordion
        var accordion1 = document.createElement("div");
        //div setAttribute
        accordion1.setAttribute("class", "accordion");
        accordion1.setAttribute("id", "accordion1");
        //accordion1.setAttribute("role", "tablist");
        //accordion1.setAttribute("aria-multiselectable", "true");
        form.appendChild(accordion1);
  
  
        //divTag for accordion
        var card = document.createElement("div");
        //div setAttribute
        card.setAttribute("class", "card");
        accordion1.appendChild(card);
  
        //divTag for accordion
        var cardHeader = document.createElement("div");
        //div setAttribute
        cardHeader.setAttribute("class", "card-header");
        cardHeader.setAttribute("role", "tab");
        cardHeader.setAttribute("id", "headingOne");
        card.appendChild(cardHeader);
  
        //h5Tag for accordion
        var groupHeader = document.createElement("h5");
        //div setAttribute
        groupHeader.setAttribute("class", "mb-0");
        cardHeader.appendChild(groupHeader);
  
        //aTag for accordion
        var groupText = document.createElement("a");
        //a setAttribute
        groupText.setAttribute("class", "ext-body d-block p-3 m-n3 link_color");
        groupText.setAttribute("data-toggle", "collapse");
        groupText.setAttribute("href", "#collapseOne");
        groupText.setAttribute("role", "button");
        groupText.setAttribute("aria-expanded", "false");
        groupText.setAttribute("aria-controls", "collapseOne");
        groupText.innerHTML = "グループ";
        groupHeader.appendChild(groupText);
  
        // group count Loop
        var groupInfo1 = document.createElement("div");
        //a setAttribute
        groupInfo1.setAttribute("id", "collapseOne");
        groupInfo1.setAttribute("class", "collapse");
        groupInfo1.setAttribute("role", "tabpanel");
        groupInfo1.setAttribute("aria-labelledby", "headingOne");
        groupInfo1.setAttribute("data-parent", "#accordion1");
        card.appendChild(groupInfo1);
  
        var groupData = {
          // n人分
          list: [
            { id: '1', name: 'TK-corporation-1', icon: 'img/pichu.png' },
            { id: '2', name: 'TK-corporation-2', icon: 'img/pikachu.jpg' },
            { id: '3', name: 'TK-corporation-3', icon: 'img/raichu.png' },
          ]
        };
        //loop
        for (var i = 0; i < groupData.list.length; i++) {
          var groupInfo2 = document.createElement("div");
          groupInfo2.setAttribute("class", "card-body");
          //image
          var imgIcon = document.createElement("img");
          imgIcon.setAttribute("src", groupData.list[i].icon);
          imgIcon.setAttribute("alt", "image");
          imgIcon.setAttribute("class", "rounded-circle icon");
          imgIcon.style.marginRight = "10px";
          groupInfo2.appendChild(imgIcon);
          //name
          var nextLink = document.createElement("a");
          nextLink.setAttribute("id", "nextLink" + i);
          nextLink.setAttribute("href", "./talkRoom.html");
          nextLink.setAttribute("class", "link_color");
          nextLink.innerHTML = groupData.list[i].name;
          groupInfo2.appendChild(nextLink);
          groupInfo1.appendChild(groupInfo2);
        }
  
  
        //divTag for accordion
        var accordion2 = document.createElement("div");
        //div setAttribute
        accordion2.setAttribute("class", "accordion");
        accordion2.setAttribute("id", "accordion2");
        accordion2.setAttribute("role", "tablist");
        accordion2.setAttribute("aria-multiselectable", "true");
        form.appendChild(accordion2);
  
        //card tag
        //divTag for accordion
        var card2 = document.createElement("div");
        //div setAttribute
        card.setAttribute("class", "card");
        accordion2.appendChild(card2);
  
        //divTag for accordion
        var cardHeader2 = document.createElement("div");
        //div setAttribute
        cardHeader2.setAttribute("class", "card-header");
        cardHeader2.setAttribute("role", "tab");
        cardHeader2.setAttribute("id", "headingTwo");
        card2.appendChild(cardHeader2);
  
        //h5Tag for accordion
        var friendHeader = document.createElement("h5");
        //div setAttribute
        friendHeader.setAttribute("class", "mb-0");
        cardHeader2.appendChild(friendHeader);
  
        //aTag for accordion
        var friendText = document.createElement("a");
        //a setAttribute
        friendText.setAttribute("class", "ext-body d-block p-3 m-n3 link_color");
        friendText.setAttribute("data-toggle", "collapse");
        friendText.setAttribute("href", "#collapseTwo");
        friendText.setAttribute("role", "button");
        friendText.setAttribute("aria-expanded", "false");
        friendText.setAttribute("aria-controls", "collapseTwo");
        friendText.innerHTML = "友達";
        friendHeader.appendChild(friendText);
  
        // friend count Loop
        var friendInfo1 = document.createElement("div");
        //a setAttribute
        friendInfo1.setAttribute("id", "collapseTwo");
        friendInfo1.setAttribute("class", "collapse");
        friendInfo1.setAttribute("role", "tabpanel");
        friendInfo1.setAttribute("aria-labelledby", "headingTwo");
        friendInfo1.setAttribute("data-parent", "#accordion2");
        card2.appendChild(friendInfo1);
  
        var friendData = {
          // n人分
          list: [
            { id: '1', name: '山田', icon: 'img/myface.png' },
            { id: '2', name: '山田 二郎', icon: 'img/myface.png' },
            { id: '3', name: '山田 三郎', icon: 'img/myface.png' },
            { id: '4', name: '山田 4郎', icon: 'img/myface.png' },
            { id: '5', name: '山田 5郎', icon: 'img/myface.png' },
            { id: '6', name: '山田 6郎', icon: 'img/myface.png' },
          ]
        };
        //loop
        for (var i = 0; i < friendData.list.length; i++) {
  
          // group count Loop
          var friendInfo2 = document.createElement("div");
          //a setAttribute
          friendInfo2.setAttribute("class", "card-body");
  
          //image
          var imgIcon = document.createElement("img");
          imgIcon.setAttribute("src", friendData.list[i].icon);
          imgIcon.setAttribute("alt", "image");
          imgIcon.setAttribute("class", "rounded-circle icon");
          imgIcon.style.marginRight = "10px";
          friendInfo2.appendChild(imgIcon);
  
          //???
          var nextLink = document.createElement("a");
          nextLink.setAttribute("id", "nextLink" + i);
          nextLink.setAttribute("href", "./talkRoom.html");
          nextLink.setAttribute("class", "link_color");
          nextLink.innerHTML = friendData.list[i].name;
          // var friendName = document.createElement("input");
          // friendName.setAttribute("style", "border: none; outline: none;");
          // friendName.setAttribute("value", friendData.list[i].name);
  
          //document.getElementById("nextLink1").innerText=friendData.list[i].name;
          friendInfo2.appendChild(nextLink);
          //nextLink.appendChild(friendName);
          friendInfo1.appendChild(friendInfo2);
        }
  
        var buttonArea = document.createElement("div");
        buttonArea.setAttribute("class", "btn_position");
        form.appendChild(buttonArea);
  
        var button1 = document.createElement("button");
        button1.setAttribute("type", "submit");
        button1.setAttribute("class", "btn btn-secondary");
        button1.innerHTML = "ログアウト";
        buttonArea.appendChild(button1);
      }
      console.log(screen.width);