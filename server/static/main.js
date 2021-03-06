const form = document.querySelector("form");
form.addEventListener("submit", handleSubmit);

const state = document.querySelector("#state");
state.addEventListener("click", handleStateChange);
handleStateChange();

async function handleSubmit(e) {
    e.preventDefault();

    const formData = new FormData(e.target);
    const obj = Object.fromEntries(formData.entries());
    form.reset();

    const state = getState();
    const url = state === "encrypt" ? "/encrypt" : "/decrypt";
    const response = await fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(obj),
    });

    if (response.status === 200) {
        const json = await response.json();
        renderResponse(json.plainText);
    } else if (response.status === 201) {
        renderResponse("Success!");
    } else {
        const text = await response.text();
        renderResponse(text);
    }
}

function renderResponse(str) {
    clearResponseIfPresent();
    const resp = document.createElement("summary");
    resp.setAttribute("id", "response");
    resp.textContent = str;

    const body = document.querySelector("body");
    body.append(resp);
    resp.addEventListener("click", function (e) {
        copyToClipboard(e.target.textContent);
    });
}

function getState() {
    return document.querySelector("input[name=state]:checked").value;
}

function handleStateChange() {
    const dataInput = document.querySelector("#data-input");
    const state = getState();
    if (state === "decrypt") {
        dataInput.style.visibility = "hidden";
    } else {
        dataInput.style.visibility = "visible";
    }
    clearResponseIfPresent();
}

function copyToClipboard(str) {
    const el = document.createElement("textarea");
    el.value = str;
    document.body.appendChild(el);
    el.select();
    document.execCommand("copy");
    document.body.removeChild(el);
}

function clearResponseIfPresent() {
    const r = document.querySelector("#response");
    if (r) {
        r.remove();
    }
}
