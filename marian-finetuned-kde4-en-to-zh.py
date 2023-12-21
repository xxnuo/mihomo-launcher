# uvicorn main:app --port 7777 --reload

from transformers import MarianMTModel, MarianTokenizer, pipeline
import torch
from fastapi import FastAPI
from pydantic import BaseModel
from typing import List
import os

model_name = "Normal1919/Marian-NMT-en-zh-lil-fine-tune"

device = torch.device("cuda" if torch.cuda.is_available() else "cpu")
tokenizer = MarianTokenizer.from_pretrained(model_name)#.to(device)
model = MarianMTModel.from_pretrained(model_name).to(device)
pipe = pipeline("translation", model=model, tokenizer=tokenizer, device=0)

app = FastAPI()

class TranslationRequest(BaseModel):
    source_lang: str
    target_lang: str
    text_list: List[str]

class TranslationResponse(BaseModel):
    detected_source_lang: str
    text: str

@app.post("/translate")
async def translate_text(request: TranslationRequest):
    # 在此处编写翻译逻辑，将源语言的文本翻译为目标语言
    # print(request.text_list)
    # input()
    model_ret = pipe(request.text_list)
    
    for ret in model_ret:
        ret.update({"detected_source_lang": request.source_lang})
    print(model_ret)
    return {"translations": model_ret}
    # input()
    # model_ret = [tokenizer.decode(t, skip_special_tokens=True) for t in model_ret]
    
    # translations = []

    # for text in model_ret:
    #     ret = TranslationResponse(detected_source_lang=request.source_lang, text=text)
    #     translations.append(ret)

    # return {"translations": translations}

    # 在此处编写翻译逻辑，将源语言的文本翻译为目标语言
#     translations = []

#     for text in request.text_list:
#         # 在此处调用翻译服务或函数，将源语言的文本翻译为目标语言
#         translated_text = translate(request.source_lang, request.target_lang, text)
#         translation = TranslationResponse(detected_source_lang=request.source_lang, text=translated_text)
#         translations.append(translation)

#     return {"translations": translations}    

# def translate(source_lang: str, target_lang: str, text: str) -> str:
#     # 在此处编写调用翻译服务或函数的代码，将源语言的文本翻译为目标语言
#     # 返回翻译后的文本
#     return text


# pipe = pipeline("translation", model="peteryushunli/marian-finetuned-kde4-en-to-zh")

# while True:
#     input_str = input("Enter a sentence to translate: ")
#     print(pipe(input_str)[0]['translation_text'])
    
