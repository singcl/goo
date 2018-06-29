const readline = require('readline');
const fs = require('fs');

const rl = readline.createInterface({
    input: fs.createReadStream('README.md'),
    output: null,
    crlfDelay: Infinity
});

rl.on('line', (line) => {
    console.log(`文件的单行内容：${line}`);
});
