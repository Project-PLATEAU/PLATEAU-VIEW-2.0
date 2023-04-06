import fs from "fs";
import path from "path";

import archiver from "archiver";
import Yml from "yml";

const root = path.resolve("./");

const pluginInfo = Yml.load(`${root}/public/reearth.yml`);
const filename = `${pluginInfo.id}-${pluginInfo.version}.zip`;

const output = fs.createWriteStream(`${root}/dist/${filename}`);
const archive = archiver("zip", { zlib: { level: 9 } });

archive.pipe(output);

archive.directory(`${root}/dist/plugin/`, false);

archive.finalize();
