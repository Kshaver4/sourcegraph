import { CodeHosts, RepogroupMetadata } from './types'
import { SearchPatternType } from '../../../shared/src/graphql/schema'
import * as React from 'react'

export const python2To3Metadata: RepogroupMetadata = {
    title: 'Refactor Python 2 to 3',
    name: 'python',
    url: '/refactor-python2-to-3',
    repositories: [
        { name: 'donnemartin/system-design-primer', codehost: CodeHosts.GITHUB },
        { name: 'public-apis/public-apis', codehost: CodeHosts.GITHUB },
        { name: 'TheAlgorithms/Python', codehost: CodeHosts.GITHUB },
        { name: 'ytdl-org/youtube-dl', codehost: CodeHosts.GITHUB },
        { name: 'tensorflow/models', codehost: CodeHosts.GITHUB },
        { name: 'nvbn/thefuck', codehost: CodeHosts.GITHUB },
        { name: 'pallets/flask', codehost: CodeHosts.GITHUB },
        { name: 'django/django', codehost: CodeHosts.GITHUB },
        { name: 'keras-team/keras', codehost: CodeHosts.GITHUB },
        { name: 'jakubroztocil/httpie', codehost: CodeHosts.GITHUB },
        { name: 'josephmisiti/awesome-machine-learning', codehost: CodeHosts.GITHUB },
        { name: 'ansible/ansible', codehost: CodeHosts.GITHUB },
        { name: 'psf/requests', codehost: CodeHosts.GITHUB },
        { name: 'scikit-learn/scikit-learn', codehost: CodeHosts.GITHUB },
        { name: 'scrapy/scrapy', codehost: CodeHosts.GITHUB },
        { name: 'minimaxir/big-list-of-naughty-strings', codehost: CodeHosts.GITHUB },
        { name: 'ageitgey/face_recognition', codehost: CodeHosts.GITHUB },
        { name: 'home-assistant/core', codehost: CodeHosts.GITHUB },
        { name: 'soimort/you-get', codehost: CodeHosts.GITHUB },
        { name: 'shadowsocks/shadowsocks', codehost: CodeHosts.GITHUB },
        { name: 'python/cpython', codehost: CodeHosts.GITHUB },
        { name: 'deepfakes/faceswap', codehost: CodeHosts.GITHUB },
        { name: '521xueweihan/HelloGitHub', codehost: CodeHosts.GITHUB },
        { name: 'apache/incubator-superset', codehost: CodeHosts.GITHUB },
        { name: 'Avik-Jain/100-Days-Of-ML-Code', codehost: CodeHosts.GITHUB },
        { name: 'testerSunshine/12306', codehost: CodeHosts.GITHUB },
        { name: 'isocpp/CppCoreGuidelines', codehost: CodeHosts.GITHUB },
        { name: '0voice/interview_internal_reference', codehost: CodeHosts.GITHUB },
        { name: 'huggingface/transformers', codehost: CodeHosts.GITHUB },
        { name: 'floodsung/Deep-Learning-Papers-Reading-Roadmap', codehost: CodeHosts.GITHUB },
        { name: 'pandas-dev/pandas', codehost: CodeHosts.GITHUB },
        { name: 'apachecn/AiLearning', codehost: CodeHosts.GITHUB },
        { name: 'getsentry/sentry', codehost: CodeHosts.GITHUB },
        { name: 'faif/python-patterns', codehost: CodeHosts.GITHUB },
        { name: 'localstack/localstack', codehost: CodeHosts.GITHUB },
        { name: 'facebookresearch/Detectron', codehost: CodeHosts.GITHUB },
        { name: 'google-research/bert', codehost: CodeHosts.GITHUB },
        { name: 'fxsjy/jieba', codehost: CodeHosts.GITHUB },
        { name: '3b1b/manim', codehost: CodeHosts.GITHUB },
        { name: '0xAX/linux-insides', codehost: CodeHosts.GITHUB },
        { name: 'satwikkansal/wtfpython', codehost: CodeHosts.GITHUB },
        { name: 'openai/gym', codehost: CodeHosts.GITHUB },
        { name: 'fighting41love/funNLP', codehost: CodeHosts.GITHUB },
        { name: 'pypa/pipenv', codehost: CodeHosts.GITHUB },
        { name: 'donnemartin/interactive-coding-challenges', codehost: CodeHosts.GITHUB },
        { name: 'tornadoweb/tornado', codehost: CodeHosts.GITHUB },
        { name: 'mitmproxy/mitmproxy', codehost: CodeHosts.GITHUB },
        { name: 'chubin/cheat.sh', codehost: CodeHosts.GITHUB },
        { name: 'donnemartin/data-science-ipython-notebooks', codehost: CodeHosts.GITHUB },
        { name: 'trailofbits/algo', codehost: CodeHosts.GITHUB },
        { name: 'encode/django-rest-framework', codehost: CodeHosts.GITHUB },
        { name: 'CorentinJ/Real-Time-Voice-Cloning', codehost: CodeHosts.GITHUB },
        { name: 'keon/algorithms', codehost: CodeHosts.GITHUB },
        { name: 'geekcomputers/Python', codehost: CodeHosts.GITHUB },
        { name: 'd2l-ai/d2l-zh', codehost: CodeHosts.GITHUB },
        { name: 'google/python-fire', codehost: CodeHosts.GITHUB },
        { name: 'matterport/Mask_RCNN', codehost: CodeHosts.GITHUB },
        { name: 'apache/airflow', codehost: CodeHosts.GITHUB },
        { name: 'yunjey/pytorch-tutorial', codehost: CodeHosts.GITHUB },
        { name: 'explosion/spaCy', codehost: CodeHosts.GITHUB },
        { name: 'eriklindernoren/ML-From-Scratch', codehost: CodeHosts.GITHUB },
        { name: 'psf/black', codehost: CodeHosts.GITHUB },
        { name: 'nicolargo/glances', codehost: CodeHosts.GITHUB },
        { name: 'drduh/macOS-Security-and-Privacy-Guide', codehost: CodeHosts.GITHUB },
        { name: 'HelloZeroNet/ZeroNet', codehost: CodeHosts.GITHUB },
        { name: 'sebastianruder/NLP-progress', codehost: CodeHosts.GITHUB },
        { name: 'reddit-archive/reddit', codehost: CodeHosts.GITHUB },
        { name: 'wangzheng0822/algo', codehost: CodeHosts.GITHUB },
        { name: 'tensorflow/magenta', codehost: CodeHosts.GITHUB },
        { name: 'celery/celery', codehost: CodeHosts.GITHUB },
        { name: 'swisskyrepo/PayloadsAllTheThings', codehost: CodeHosts.GITHUB },
        { name: 'tqdm/tqdm', codehost: CodeHosts.GITHUB },
        { name: 'machinelearningmindset/TensorFlow-Course', codehost: CodeHosts.GITHUB },
        { name: 'binux/pyspider', codehost: CodeHosts.GITHUB },
        { name: 'ipython/ipython', codehost: CodeHosts.GITHUB },
        { name: 'iperov/DeepFaceLab', codehost: CodeHosts.GITHUB },
        { name: 'luong-komorebi/Awesome-Linux-Software', codehost: CodeHosts.GITHUB },
        { name: 'wangshub/wechat_jump_game', codehost: CodeHosts.GITHUB },
        { name: 'huge-success/sanic', codehost: CodeHosts.GITHUB },
        { name: 'tiangolo/fastapi', codehost: CodeHosts.GITHUB },
        { name: 'cool-RR/PySnooper', codehost: CodeHosts.GITHUB },
        { name: 'bokeh/bokeh', codehost: CodeHosts.GITHUB },
        { name: 'spotify/luigi', codehost: CodeHosts.GITHUB },
        { name: 'bitcoinbook/bitcoinbook', codehost: CodeHosts.GITHUB },
        { name: 'nginx-proxy/nginx-proxy', codehost: CodeHosts.GITHUB },
        { name: 'pytorch/examples', codehost: CodeHosts.GITHUB },
        { name: 'locustio/locust', codehost: CodeHosts.GITHUB },
        { name: 'gto76/python-cheatsheet', codehost: CodeHosts.GITHUB },
        { name: 'OWASP/CheatSheetSeries', codehost: CodeHosts.GITHUB },
        { name: 'fabric/fabric', codehost: CodeHosts.GITHUB },
        { name: 'youfou/wxpy', codehost: CodeHosts.GITHUB },
        { name: 'cookiecutter/cookiecutter', codehost: CodeHosts.GITHUB },
        { name: 'junyanz/pytorch-CycleGAN-and-pix2pix', codehost: CodeHosts.GITHUB },
        { name: 'openai/gpt-2', codehost: CodeHosts.GITHUB },
        { name: 'Kr1s77/awesome-python-login-model', codehost: CodeHosts.GITHUB },
        { name: 'powerline/powerline', codehost: CodeHosts.GITHUB },
        { name: 'sherlock-project/sherlock', codehost: CodeHosts.GITHUB },
        { name: 'matplotlib/matplotlib', codehost: CodeHosts.GITHUB },
        { name: 'microsoft/cascadia-code', codehost: CodeHosts.GITHUB },
        { name: 'deezer/spleete', codehost: CodeHosts.GITHUB },
        { name: 'robcalcroft/react-use-lazy-load-image', codehost: CodeHosts.GITHUB },
        { name: 'intercaetera/react-use-message-bar', codehost: CodeHosts.GITHUB },
        { name: 'wowlusitong/react-use-modal', codehost: CodeHosts.GITHUB },
        { name: 'zhangkaiyulw/react-use-path', codehost: CodeHosts.GITHUB },
        { name: 'neo/react-use-scroll-position', codehost: CodeHosts.GITHUB },
        { name: 'lessmess-dev/react-use-trigger', codehost: CodeHosts.GITHUB },
        { name: 'perlin-network/react-use-wavelet', codehost: CodeHosts.GITHUB },
        { name: 'streamich/react-use', codehost: CodeHosts.GITHUB },
        { name: 'GeDiez/react-use-formless', codehost: CodeHosts.GITHUB },
        { name: 'venil7/react-usemiddleware', codehost: CodeHosts.GITHUB },
        { name: 'alex-cory/react-useportal', codehost: CodeHosts.GITHUB },
        { name: 'vardius/react-user-media', codehost: CodeHosts.GITHUB },
        { name: 'f/react-wait', codehost: CodeHosts.GITHUB },
        { name: 'AvraamMavridis/react-window-communication-hook', codehost: CodeHosts.GITHUB },
        { name: 'yesmeck/react-with-hooks', codehost: CodeHosts.GITHUB },
        { name: 'mfrachet/reaktion', codehost: CodeHosts.GITHUB },
        { name: 'iusehooks/redhooks', codehost: CodeHosts.GITHUB },
        { name: 'ianobermiller/redux-react-hook', codehost: CodeHosts.GITHUB },
        { name: 'regionjs/region-core', codehost: CodeHosts.GITHUB },
        { name: 'imbhargav5/rehooks-visibility-sensor', codehost: CodeHosts.GITHUB },
        { name: 'pedronasser/resynced', codehost: CodeHosts.GITHUB },
        { name: 'brn/rrh', codehost: CodeHosts.GITHUB },
        { name: 'LeetCode-OpenSource/rxjs-hooks', codehost: CodeHosts.GITHUB },
        { name: 'dejorrit/scroll-data-hook', codehost: CodeHosts.GITHUB },
        { name: 'style-hook/style-hook', codehost: CodeHosts.GITHUB },
        { name: 'vercel/swr', codehost: CodeHosts.GITHUB },
        { name: 'jaredpalmer/the-platform', codehost: CodeHosts.GITHUB },
        { name: 'danieldelcore/trousers', codehost: CodeHosts.GITHUB },
        { name: 'mauricedb/use-abortable-fetch', codehost: CodeHosts.GITHUB },
        { name: 'awmleer/use-action', codehost: CodeHosts.GITHUB },
        { name: 'awmleer/use-async-memo', codehost: CodeHosts.GITHUB },
        { name: 'lowewenzel/use-autocomplete', codehost: CodeHosts.GITHUB },
        { name: 'sergey-s/use-axios-react', codehost: CodeHosts.GITHUB },
        { name: 'zcallan/use-browser-history', codehost: CodeHosts.GITHUB },
        { name: 'samjbmason/use-cart', codehost: CodeHosts.GITHUB },
        { name: 'CharlesStover/use-clippy', codehost: CodeHosts.GITHUB },
        { name: 'dai-shi/use-context-selector', codehost: CodeHosts.GITHUB },
        { name: 'oktaysenkan/use-countries', codehost: CodeHosts.GITHUB },
        { name: 'xnimorz/use-debounce', codehost: CodeHosts.GITHUB },
        { name: 'sandiiarov/use-deep-compare', codehost: CodeHosts.GITHUB },
        { name: 'kentcdodds/use-deep-compare-effect', codehost: CodeHosts.GITHUB },
        { name: 'gregnb/use-detect-print', codehost: CodeHosts.GITHUB },
        { name: 'CharlesStover/use-dimensions', codehost: CodeHosts.GITHUB },
        { name: 'zattoo/use-double-click', codehost: CodeHosts.GITHUB },
        { name: 'sandiiarov/use-events', codehost: CodeHosts.GITHUB },
        { name: 'CharlesStover/use-force-update', codehost: CodeHosts.GITHUB },
        { name: 'sandiiarov/use-hotkeys', codehost: CodeHosts.GITHUB },
        { name: 'therealparmesh/use-hovering', codehost: CodeHosts.GITHUB },
        { name: 'ava/use-http', codehost: CodeHosts.GITHUB },
        { name: 'immerjs/use-immer', codehost: CodeHosts.GITHUB },
        { name: 'immerjs/immer', codehost: CodeHosts.GITHUB },
        { name: 'neighborhood999/use-input-file', codehost: CodeHosts.GITHUB },
        { name: 'helderburato/use-is-mounted-ref', codehost: CodeHosts.GITHUB },
        { name: 'davidicus/use-lang-direction', codehost: CodeHosts.GITHUB },
        { name: 'streamich/use-media', codehost: CodeHosts.GITHUB },
        { name: 'dimitrinicolas/use-mouse-action', codehost: CodeHosts.GITHUB },
        { name: 'jschloer/use-multiselect', codehost: CodeHosts.GITHUB },
        { name: 'wellyshen/use-places-autocomplete', codehost: CodeHosts.GITHUB },
        { name: 'sandiiarov/use-popper', codehost: CodeHosts.GITHUB },
        { name: 'alex-cory/use-react-modal', codehost: CodeHosts.GITHUB },
        { name: 'CharlesStover/use-react-router', codehost: CodeHosts.GITHUB },
        { name: 'tedstoychev/use-reactive-state', codehost: CodeHosts.GITHUB },
        { name: 'dai-shi/use-reducer-async', codehost: CodeHosts.GITHUB },
        { name: 'flepretre/use-redux', codehost: CodeHosts.GITHUB },
        { name: 'tudorgergely/use-scroll-to-bottom', codehost: CodeHosts.GITHUB },
        { name: 'sandiiarov/use-simple-undo', codehost: CodeHosts.GITHUB },
        { name: 'mfrachet/use-socketio', codehost: CodeHosts.GITHUB },
        { name: 'iamgyz/use-socket.io-client', codehost: CodeHosts.GITHUB },
        { name: 'kmoskwiak/useSSE', codehost: CodeHosts.GITHUB },
        { name: 'alex-cory/use-ssr', codehost: CodeHosts.GITHUB },
        { name: 'haydn/use-state-snapshots', codehost: CodeHosts.GITHUB },
        { name: 'philipp-spiess/use-substate', codehost: CodeHosts.GITHUB },
        { name: 'octet-stream/use-suspender', codehost: CodeHosts.GITHUB },
        { name: 'streamich/use-t', codehost: CodeHosts.GITHUB },
        { name: 'homerchen19/use-undo', codehost: CodeHosts.GITHUB },
        { name: 'donavon/use-dark-mode', codehost: CodeHosts.GITHUB },
        { name: 'KATT/use-is-typing', codehost: CodeHosts.GITHUB },
        { name: 'pranesh239/use-key-capture', codehost: CodeHosts.GITHUB },
        { name: 'tranbathanhtung/usePosition', codehost: CodeHosts.GITHUB },
        { name: 'Tweries/useReducerWithLocalStorage', codehost: CodeHosts.GITHUB },
        { name: 'pankod/react-hooks-screen-type', codehost: CodeHosts.GITHUB },
        { name: 'Purii/react-use-scrollspy', codehost: CodeHosts.GITHUB },
        { name: 'JCofman/react-hook-use-service-worker', codehost: CodeHosts.GITHUB },
        { name: 'bboydflo/use-value-after', codehost: CodeHosts.GITHUB },
        { name: 'wednesday-solutions/react-screentype-hoo', codehost: CodeHosts.GITHUB },
    ],
    description: `This repository group contains python 2 and 3 repositories and corresponding search examples.
                        The search examples will help you find code you need to refactor and review for migrating from
                        Python 2 to Python 3.`,
    examples: [
        {
            title: 'Python 2 relative imports',
            description: 'Python 2 import style',
            exampleQuery: <>"from :[package.] import :[function.]”</>,
            rawQuery: '"from :[package.] import :[function.]”',
            patternType: SearchPatternType.structural,
        },
        {
            title: 'Python 3 explicit imports',
            description: 'Python 3 import style',
            exampleQuery: <>from B.w+ import w+</>,
            rawQuery: 'from B.w+ import w+',
            patternType: SearchPatternType.regexp,
        },
        {
            title: 'Python 2 print statements',
            exampleQuery: <>\'print ":[string]"\'</>,
            rawQuery: '\'print ":[string]"\'',
            patternType: SearchPatternType.structural,
        },
        {
            title: 'Python 3 print statements',
            exampleQuery: <>\'print ":[string]"\'</>,
            rawQuery: '\'print ":[string]"\'',
            patternType: SearchPatternType.regexp,
        },
        {
            title: 'Python 2 integer conversion',
            exampleQuery: <>float(:[arg]) / float(:[arg])</>,
            rawQuery: 'float(:[arg]) / float(:[arg])',
            patternType: SearchPatternType.structural,
        },
        {
            title: 'Python 3 integer conversion',
            exampleQuery: (
                <>
                    <span className="repogroup-page__keyword-text">lang:</span>python \\sint\\(-*\\d+\\)
                </>
            ),
            rawQuery: 'lang:python \\sint\\(-*\\d+\\)',
            patternType: SearchPatternType.regexp,
        },
    ],
    homepageDescription: 'Searches and suggestions to help migrate your code to Python 3.',
    homepageIcon:
        'data:image/png;base64,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',
}
