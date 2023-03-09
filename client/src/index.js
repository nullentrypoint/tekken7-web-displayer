import LiveGollection from "../node_modules/livegollection-client/dist/index.js";

let liveGoll = null;
let divConsole = null;

function templateCard(message) {
    return `
    <div class="card">
        <div class="avatar">
            <a href="`+ profileUrl(message.SteamID) +`" target="_blank">
                <img class="avatar_img" src="`+ message.AvatarUrl +`">
            </a>
        </div>
        <div>
            <div class="persona_name">
                <span>`+ escapeHTML(message.Name) +`</span>
            </div>
            <div class="location">`
            +(message.CountryCode != '' ? `<img class="profile_flag" src="https://community.cloudflare.steamstatic.com/public/images/countryflags/`+ message.CountryCode +`.gif">` : '')+`
                
                <span>`+ escapeHTML(message.Location) +`</span>
            </div>
            <div class="ip">
                <span>IP: `+ escapeHTML(message.IP) +` </span>
            </div>
        </div>
        <div class="message_time">
                <span>`+ escapeHTML(message.Time) +`</span>
        </div>
    </div>
    `
}

function profileUrl(steamID) {
    return "https://steamcommunity.com/profiles/" + steamID + "/"
}

function profileLink(steamID, title) {
    return '<a href="' + profileUrl(steamID) + '" target="_blank">' + escapeHTML(title) + '</a>'
}

function addMessageToInbox(message) {
    if (message.info === undefined) {
        return
    }

    const newLine = document.createElement("li");
    newLine.id = message.id;
    newLine.innerHTML = templateCard(message.info);
    divConsole.insertBefore(newLine, divConsole.firstChild)
}

window.onload = () => {
    const host = window.location.host;
    liveGoll = new LiveGollection("ws://"+ host +"/livegollection");
    divConsole = document.getElementById("console");

    liveGoll.oncreate = (message) => {
        addMessageToInbox(message);
    };

    liveGoll.onupdate = (message) => {
        // Modify element in the DOM
        //addMessageToInbox(message);
    };
    
    liveGoll.ondelete = (item) => {
        // Delete item from the DOM
    };
};

const escapeHTML = str =>
  str.replace(
    /[&<>'"]/g,
    tag =>
      ({
        '&': '&amp;',
        '<': '&lt;',
        '>': '&gt;',
        "'": '&#39;',
        '"': '&quot;'
      }[tag] || tag)
  );