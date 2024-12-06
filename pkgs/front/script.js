/**
 * データベースサーバへGETリクエストを投げるためのヘルパー関数。
 *
 * エラーが発生した場合、consoleにerror出力し、nullを返す。
 */
async function get(url, isDebug) {
  const base = isDebug ? "http://localhost:63245" : "http://skd-sv.skdassoc.work"
  try {
    const res = await fetch(base + url, {
      method: "GET",
      headers: {
          "Content-Type": "application/json",
      },
    })
    if (!res.ok) {
      console.error(res)
      return null
    }

    return await res.json()
  } catch (e) {
    console.error(e)
    return null
  }
}

/**
 * データベースサーバへPOSTリクエストを投げるためのヘルパー関数。
 *
 * エラーが発生した場合、consoleにerror出力し、nullを返す。
 */
async function post(url, isDebug, body) {
  const base = isDebug ? "http://localhost:63245" : "http://skd-sv.skdassoc.work"
  try {
    const res = await fetch(base + url, {
      method: "POST",
      headers: {
          "Content-Type": "application/json",
      },
      body: JSON.stringify(body),
    })
    if (!res.ok) {
      console.error(res)
      return null
    }

    return await res.json()
  } catch (e) {
    console.error(e)
    return null
  }
}
