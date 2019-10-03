package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	DimIn     = 4    // 输入样本的维数
	DimOut    = 1    // 输出样本的维数
	HidePoint = 8    // 隐藏节点
	MaxIter   = 1000 // 最大迭代次数
	Alpha     = 0.35 // 学习率
)

var counterResults [7]int
var yns = make([]float64, 0)

//获取训练样本
func getInputTrain(filename string) [][]float64 {
	flysnowRegexp := regexp.MustCompile(`[0-9]*\.?[0-9]+`)
	fp, err := os.Open(filename)
	defer fp.Close()
	if err != nil {
		fmt.Println("Open input train error",err)
		os.Exit(0)
	}
	content, _ := ioutil.ReadAll(fp)
	//fmt.Println(content)
	s_content := string(content)
	lines := strings.Split(s_content, "\n")
	ret := make([][]float64, 0)
	for _, line := range lines {

		line = strings.TrimRight(line, "\r\n")

		if len(line) == 0 {
			continue
		}
		tup := flysnowRegexp.FindAll([]byte(line),-1)
		X := make([]float64, 0)

		for _, x := range tup {

			f_x, _ := strconv.ParseFloat(string(x), 64)
			X = append(X, f_x)
		}
		ret = append(ret, X)

	}
	return ret
}

//获取期望样本
func getOutputTrain(filename string) [][]float64 {
	flysnowRegexp := regexp.MustCompile(`[0-9]*\.?[0-9]+`)
	fp, err := os.Open(filename)
	defer fp.Close()
	if err != nil {
		fmt.Println("Open input train error")
		os.Exit(0)
	}
	content, _ := ioutil.ReadAll(fp)
	//fmt.Println(content)
	s_content := string(content)
	lines := strings.Split(s_content, "\n")
	ret := make([][]float64, 0)
	for _, line := range lines {
		line = strings.TrimRight(line, "\r\n")
		if len(line) == 0 {
			continue
		}
		tup := flysnowRegexp.FindAll([]byte(line),-1)
		X := make([]float64, 0)

		for _, x := range tup {

			f_x, _ := strconv.ParseFloat(string(x), 64)
			X = append(X, f_x)
		}
		ret = append(ret, X)
	}
	return ret
}

//训练样本归一化
func inputNormalization(inputTrain [][]float64) [][]float64 {
	maxInputTrain := [DimIn]float64{}
	minInputTrain := [DimIn]float64{}
	for i := 0; i < DimIn; i++ {
		maxInputTrain[i], minInputTrain[i] = inputTrain[0][i], inputTrain[0][i]
		for j := 1; j < len(inputTrain); j++ {
			if maxInputTrain[i] < inputTrain[j][i] {
				maxInputTrain[i] = inputTrain[j][i]
			}
			if minInputTrain[i] > inputTrain[j][i] {
				minInputTrain[i] = inputTrain[j][i]
			}
		}
	}

	//归一化
	ret := make([][]float64, 0)
	for i := 0; i < len(inputTrain); i++ {
		temp := make([]float64, 0)
		for j := 0; j < DimIn; j++ {
			y := (0.02 + 0.996*(inputTrain[i][j]-minInputTrain[j])) / (maxInputTrain[j] - minInputTrain[j])
			temp = append(temp, y)
		}

		ret = append(ret, temp)
	}

	return ret
}

//输出期望归一化
func outputNormalization(outputTrain [][]float64) [][]float64 {
	maxOutputTrain := [DimOut]float64{}
	minOutputTrain := [DimOut]float64{}
	for i := 0; i < DimOut; i++ {
		maxOutputTrain[i], minOutputTrain[i] = outputTrain[0][i], outputTrain[0][i]
		for j := 1; j < len(outputTrain); j++ {
			if maxOutputTrain[i] < outputTrain[j][i] {
				maxOutputTrain[i] = outputTrain[j][i]
			}
			if minOutputTrain[i] > outputTrain[j][i] {
				minOutputTrain[i] = outputTrain[j][i]
			}
		}
	}

	//归一化
	ret := make([][]float64, 0)
	for i := 0; i < len(outputTrain); i++ {
		temp := make([]float64, 0)
		for j := 0; j < DimOut; j++ {
			y := (0.02 + 0.996*(outputTrain[i][j]-minOutputTrain[j])) / (maxOutputTrain[j] - minOutputTrain[j])
			temp = append(temp, y)
		}

		ret = append(ret, temp)
	}

	return ret
}

//开始训练样本
func trainNet(inputTrain, outputTrain [][]float64) {
	c := [HidePoint][DimIn]float64{}
	c_1 := [HidePoint][DimIn]float64{}
	b := [HidePoint][DimIn]float64{}
	b_1 := [HidePoint][DimIn]float64{}

	//初始化参数
	rand.Seed(time.Now().Unix())
	for i := 0; i < HidePoint; i++ {
		for j := 0; j < DimIn; j++ {
			c[i][j] = (float64)(rand.Uint32()) / (float64)(^(uint32)(0))
			b[i][j] = (float64)(rand.Uint32()) / (float64)(^(uint32)(0))
			c_1[i][j] = c[i][j]
			b_1[i][j] = b[i][j]
		}
	}

	p := [DimIn + 1][HidePoint]float64{}
	p_1 := [DimIn + 1][HidePoint]float64{}
	for i := 0; i < DimIn+1; i++ {
		for j := 0; j < HidePoint; j++ {
			p[i][j] = 0.18
			p_1[i][j] = p[i][j]
		}
	}

	u := [DimIn][HidePoint]float64{}
	w := [HidePoint]float64{}
	y := [HidePoint]float64{}
	d_p := [HidePoint]float64{}
	// 开始训练
	for iter := 0; iter < MaxIter; iter++ {
		for k := 0; k < len(inputTrain); k++ {
			//获取模糊参数
			for i := 0; i < DimIn; i++ {
				for j := 0; j < HidePoint; j++ {
					u[i][j] = math.Exp(-1 * ((inputTrain[k][i] - c[j][i]) * (inputTrain[k][i] - c[j][i])) / b[j][i]) //模糊隶属度计算
				}
			}

			var addw float64
			for i := 0; i < HidePoint; i++ {
				var mul float64 = 1
				for j := 0; j < DimIn; j++ {
					mul = mul * u[j][i]
				}
				w[i] = mul
				addw += w[i]
			}

			//计算输出
			var addyw float64
			for i := 0; i < HidePoint; i++ {
				var sumTemp float64
				for j := 0; j < DimIn; j++ {
					sumTemp += p[j][i] * inputTrain[k][j]
				}
				y[i] = sumTemp + p[DimIn][i]
				addyw += y[i] * w[i]
			}

			yn := addyw / addw //模糊输出值
			e := outputTrain[k][0] - yn

			//修正系数p
			for i := 0; i < HidePoint; i++ {
				d_p[i] = Alpha * e * w[i] / addw
			}
			for i := 0; i < HidePoint; i++ {
				for j := 0; j < DimIn; j++ {
					p[j][i] = p_1[j][i] + d_p[i]*inputTrain[k][j]
				}
				p[DimIn][i] = p_1[DimIn][i]
			}

			//修正系数b
			for i := 0; i < HidePoint; i++ {
				for j := 0; j < DimIn; j++ {
					b[i][j] = b_1[i][j] + Alpha*e*(y[i]*addw-addyw)*(inputTrain[k][j]-c[i][j])*(inputTrain[k][j]-c[i][j])*w[i]/(b[i][j]*b[i][j]*addw*addw)
				}
			}

			//修正系数c
			for i := 0; i < HidePoint; i++ {
				for j := 0; j < DimIn; j++ {
					c[i][j] = c_1[i][j] + Alpha*e*(y[i]*addw-addyw)*2*(inputTrain[k][j]-c[i][j])*w[i]/(b[i][j]*addw*addw)
				}
			}
		}
	}

	//测试网络
	for k := 0; k < 500; k++ {//数字根据输入行数改
		for i := 0; i < DimIn; i++ {
			for j := 0; j < HidePoint; j++ {
				u[i][j] = math.Exp(-1 * ((inputTrain[k][i] - c[j][i]) * (inputTrain[k][i] - c[j][i])) / b[j][i])
			}
		}

		//模糊隶属度计算
		var addw float64
		for i := 0; i < HidePoint; i++ {
			var mul float64 = 1
			for j := 0; j < DimIn; j++ {
				mul *= u[j][i]
			}
			w[i] = mul
			addw += w[i]
		}

		//计算输出
		var addyw float64
		for i := 0; i < HidePoint; i++ {
			var sumTemp float64
			for j := 0; j < DimIn; j++ {
				sumTemp = sumTemp + p[j][i]*inputTrain[k][j]
			}
			y[i] = sumTemp + p[DimIn][i]
			addyw += y[i] * w[i]
		}

		yn := addw / addyw
		//模糊输出值
		yns = append(yns, yn)
		switch  {
		case yn<=1:
			counterResults[0]+=1
		case yn>1&&yn<=2:
			counterResults[1]+=1
		case yn>2&&yn<=3:
			counterResults[2]+=1
		case yn>3&&yn<=4:
			counterResults[3]+=1
		case yn>4&&yn<=5:
			counterResults[4]+=1
		case yn>5&&yn<=6:
			counterResults[5]+=1
		}



		//fmt.Println(yn)
	}
}

var (
	maxInputTrain  = [DimIn]float64{} //输入训练样本的最大最小值
	minInputTrain  = [DimIn]float64{}
	maxOutputTrain = [DimOut]float64{} //期望样本的最大最小值
	minOutputTrain = [DimOut]float64{}
)


