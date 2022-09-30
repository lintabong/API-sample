import { initializeApp } from "https://www.gstatic.com/firebasejs/9.6.9/firebase-app.js";
import { getDatabase, ref, get, set, child, update, remove, onValue } from "https://www.gstatic.com/firebasejs/9.6.9/firebase-database.js";

const firebaseConfig = {
    apiKey: "AIzaSyCAV4IZbWdnBo__3MQFl9nfHY016sqP6hY",
    authDomain: "classmanagerapps.firebaseapp.com",
    projectId: "classmanagerapps",
    storageBucket: "classmanagerapps.appspot.com",
    messagingSenderId: "1023940208070",
    appId: "1:saded1fd6dfec7"
  };

// Initialize Firebase
      const app = initializeApp(firebaseConfig);
      const db = getDatabase(app);
      const countRef = ref(db);

      get(ref(db, "user")).then((snapshot) => {
        if (snapshot.exists()) {
          var user = snapshot.val();
          var i = 0;
          console.log(user);
          for (const num in user) {
            i++;
            var val_username = eval("user." + num + ".username");
            var val_password = eval("user." + num + ".password");
            var val_name = eval("user." + num + ".nama");
            var val_jabatan = eval("user." + num + ".jabatan");

            var para1 = document.createElement("p");
            var para2 = document.createElement("p");
            var para3 = document.createElement("p");
            var para4 = document.createElement("p");
            var para5 = document.createElement("p");
            var btnedit = document.createElement("button");
            var btndelete = document.createElement("button");
            para1.innerHTML = i;
            para2.innerHTML = val_username;
            para3.innerHTML = val_password;
            para4.innerHTML = val_name;
            para5.innerHTML = val_jabatan;
            btnedit.innerHTML = "edit user";
            btnedit.type = "button";
            btnedit.classList.add("btn");
            btnedit.classList.add("btn-primary");
            btnedit.onclick = function () {
              location.href = "user_edit.html?name=" + num;
            };

            btndelete.innerHTML = "delete";
            btndelete.type = "button";
            btndelete.classList.add("btn");
            btndelete.classList.add("btn-danger");
            btndelete.onclick = function () {
              remove(ref(db, "user/" + num.toString()));
              location.reload(true);
            };
            document.getElementById("user_id").appendChild(para1);
            document.getElementById("user_username").appendChild(para2);
            document.getElementById("user_password").appendChild(para3);
            document.getElementById("user_nama").appendChild(para4);
            document.getElementById("user_jabatan").appendChild(para5);
            document.getElementById("acc").appendChild(btnedit);
            document.getElementById("del").appendChild(btndelete);
          }
        }
      });
