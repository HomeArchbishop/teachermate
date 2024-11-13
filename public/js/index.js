import config from "../config.js"

const jsQR = window.jsQR

const msgTipper = document.getElementById('msgTipper')
const video = document.getElementById('video')
const canvasElement = document.createElement('canvas')
const ctx = canvasElement.getContext('2d')

const apiUrl = config.DEFAULT_API_URL

navigator.mediaDevices
  .getUserMedia({ video: true })
  .then(stream => {
    msgTipper.innerHTML = 'Scanning...'
    video.srcObject = stream
    video.onloadedmetadata = function (e) {
      video.play()
    }
  }, () => {
    msgTipper.innerHTML = 'Camera is not available'
  })

let raf
function tick () {
  if (video.readyState === video.HAVE_ENOUGH_DATA) {
    canvasElement.hidden = false

    canvasElement.height = video.videoHeight
    canvasElement.width = video.videoWidth
    ctx.drawImage(video, 0, 0, canvasElement.width, canvasElement.height)
    const imageData = ctx.getImageData(0, 0, canvasElement.width, canvasElement.height)
    const code = jsQR(imageData.data, imageData.width, imageData.height, {
      inversionAttempts: 'dontInvert'
    })
    if (code) {
      const link = code.data
      sendRequest(link)
    }
  }

  cancelAnimationFrame(raf)
  raf = requestAnimationFrame(tick)
}

raf = tick()

const alreadySignedLink = new Set()
const sendRequest = async (link) => {
  const lessonId = document.getElementById('lessonIdInput').value
  if (!lessonId) {
    msgTipper.innerHTML = 'Please input lesson ID'
    return
  }

  if (alreadySignedLink.has(link + lessonId)) {
    return
  }

  alreadySignedLink.add(link + lessonId)

  const response = await fetch(apiUrl + `/api/sign?lesson_id=${lessonId}&attendance=${link}`, {
    method: 'GET'
  })
  return response.json()
}
