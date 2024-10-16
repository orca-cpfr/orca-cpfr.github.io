# !/bin/bash

# clean
rm -rf public

# build app
cd dashboard
npm install
npm run build

cd .. 

# build landing page
cd landing-page
npm install
npm run build 

cd ..

# copy 
cp -r landing-page/public .
mkdir -p public/demo
cp -r dashboard/public public/demo/dashboard

# serve 
npx serve public
