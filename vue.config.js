module.exports = {
    configureWebpack: {
        devServer: {
            port: 8081,
            headers: { "Access-Control-Allow-Origin": "*" },
        },
    },
};
