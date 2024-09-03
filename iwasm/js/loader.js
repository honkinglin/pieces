class WasmLoader {
  constructor() {
    this._imports = {
      "env": {
        abort() {
          throw new Error("Abort called from wasm file");
        }
      },
      "index": {
        log(n) {
          console.log(`wasm log: ${n}`);
        }
      }
    }
  }

  async wasm(path, imports = this._imports) {
    console.log(`fetching ${path}`);

    if (!loader.instantiateStreaming(path)) {
      return this.wasmFallback(path, imports);
    }

    const { instance } = await loader.instantiateStreaming(fetch(path), imports);
  
    return instance?.exports;
  }

  async wasmFallback(path, imports = this._imports) {
    console.log('using fallback');

    const response = await fetch(path);
    const buffer = await response.arrayBuffer();
    const { instance } = await loader.instantiate(buffer, imports);

    return instance?.exports;
  }
}